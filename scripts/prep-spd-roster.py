import pandas as pd
import click

from pandas.core.frame import DataFrame


@click.command()
@click.option('-d', '--date', required=True, type=str, prompt=True, help='The date of the roster in the format YYYY-MM-DD.')
@click.option('-i', '--in-csv', required=True, type=str, prompt=True, help='Path to the source file. Can be relative to current directory.')
@click.option('-o', '--out-csv', required=True, type=str, prompt=True, help='Path to save the resulting CSV in. Can be relative to current directory.')
def main(date, in_csv, out_csv):
    # Import CSV passed to script
    df = pd.read_csv(in_csv)

    # Rename columns to better formatting
    column_names = {
        'Name': 'full_name',
        'Badge_Num': 'badge',
        'Title_Description': 'title',
        'Unit': 'unit',
        'Unit_Description': 'unit_description'
    }
    df.rename(columns=column_names, inplace=True)

    # Strip out unwanted whitespace from any string fields in the DataFrame
    df = df.applymap(lambda x: x.strip() if isinstance(x, str) else x)

    # Split the names out
    df['first_name'] = df["full_name"].str.split(",").str[1].str.strip().str.split(" ").str[0]
    df['middle_name'] = df.apply(lambda x: middle_name(x), axis=1)
    df['last_name'] = df["full_name"].str.split(",").str[0]

    # Add date column to DataFrame
    df['date'] = date

    # Write out resulting DataFrame to CSV path specified
    df.to_csv(out_csv, index=False)


def middle_name(df: DataFrame):
    name_list = df['full_name'].split(',')[1].split(' ')
    if len(name_list) > 2:
        return ' '.join(name_list[2:])
    else:
        return ''


if __name__ == "__main__":
    main()