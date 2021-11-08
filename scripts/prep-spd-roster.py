import pandas as pd
import click

from pandas import DataFrame


@click.command()
@click.option(
    "-d",
    "--date",
    required=True,
    type=str,
    prompt=True,
    help="The date of the roster in the format YYYY-MM-DD.",
)
@click.option(
    "-i",
    "--in-csv",
    required=True,
    type=str,
    prompt=True,
    help="Path to the source file. Can be relative to current directory.",
)
@click.option(
    "-o",
    "--out-csv",
    required=True,
    type=str,
    prompt=True,
    help="Path to save the resulting CSV in. Can be relative to current directory.",
)
def main(date, in_csv, out_csv):
    # Import CSV passed to script
    df = pd.read_csv(in_csv)

    # Rename columns to better formatting
    column_names = {
        "Name": "full_name",
        "Badge_Num": "badge",
        "Serial": "badge",
        "Title_Description": "title",
        "Title Description": "title",
        "Unit": "unit",
        "Unit Description": "unit_description",
        "Unit_Description": "unit_description",
    }
    df = df.rename(columns=column_names)

    # Strip out unwanted whitespace from any string fields in the DataFrame
    df = df.applymap(lambda x: x.strip() if isinstance(x, str) else x)

    # Split the names out
    names = df["full_name"].str.split(",", expand=True)
    # Last name is predictably before the comma, so we can set it outright
    df["last_name"] = names[0]
    # Assumption: first name will always be the text before the first space
    # We can set first name to be the first split based off that
    not_last_names = names[1].str.split(n=1, expand=True)
    df["first_name"] = not_last_names[0]
    # Next we have to split apart the middle name and the suffix
    # Sometimes there some weird "K_Jr" bullshit for the middle name we have to deal with
    # This creates a dataframe with 3 new columns: 1) middle bits, 2) suffix, 3) blank because of the split
    middle_bits = not_last_names[1].str.split(
        r"(?i)((?:Jr|II|III|IV)\.?$)", expand=True
    )
    # Make all suffixes upper, except for "Jr", and trim periods
    df["suffix"] = middle_bits[1].str.strip(".").str.upper().str.replace("JR", "Jr")
    # The first part of the middle bits is now the middle name
    # (but we also have to remove the "_" or "." in some cases)
    df["middle_name"] = middle_bits[0].str.strip("_").str.strip(".")

    # Add date column to DataFrame
    df["date"] = date

    # Write out resulting DataFrame to CSV path specified
    df.to_csv(out_csv, index=False)


if __name__ == "__main__":
    main()
