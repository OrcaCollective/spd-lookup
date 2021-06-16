import pandas as pd
import sys

from pandas.core.frame import DataFrame


def main():
    date = sys.argv[1]
    in_csv = sys.argv[2]
    out_csv = sys.argv[3]

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

    # Split the names out
    df['first_name'] = df.apply(lambda x: first_name(x), axis=1)
    df['middle_name'] = df.apply(lambda x: middle_name(x), axis=1)
    df['last_name'] = df.apply(lambda x: last_name(x), axis=1)

    # Add date column to DataFrame
    df['date'] = date

    # Write out resulting DataFrame to CSV path specified
    df.to_csv(out_csv, index=False)


def first_name(df: DataFrame):
    return df['full_name'].split(' ')[1].strip()


def middle_name(df: DataFrame):
    name_list = df['full_name'].split(' ')
    if len(name_list) > 2:
        return ' '.join(name_list[2:])
    else:
        return ''


def last_name(df: DataFrame):
    return df['full_name'].split(' ')[0].strip(' ,')


if __name__ == "__main__":
    main()