from pathlib import Path
import pandas as pd
import click

from pandas import DataFrame

DEFAULT_HISTORICAL = (
    Path(__file__).parents[1] / "seed" / "Seattle-WA-Police-Department_Historical.csv"
)
COLUMNS = [
    "badge",
    "full_name",
    "title",
    "unit",
    "unit_description",
    "first_name",
    "middle_name",
    "last_name",
    "date",
]


@click.command()
@click.argument("input_csvs", nargs=-1)
@click.option(
    "-h",
    "--historical-csv",
    default=str(DEFAULT_HISTORICAL),
    type=str,
    prompt=True,
    help="Path to the historical CSV. This file will be replaced with the updated version",
)
def main(input_csvs, historical_csv):
    # Read in the historical CSV
    df = pd.read_csv(historical_csv)

    for input_csv in input_csvs:
        print(f"Adding {input_csv}")
        # Read the CSV, gathering only a subset of the columns
        xdf = pd.read_csv(input_csv, usecols=COLUMNS)
        # Append the dataframe
        df = df.append(xdf, ignore_index=True)

    # Write out resulting DataFrame to CSV path specified
    df.to_csv(historical_csv, index=False)


if __name__ == "__main__":
    main()
