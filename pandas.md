[python](./python.md)

# Pandas

> python library for data manipulation and analysis

- features
  - load and clean data
  - explore, filter, sort and group data
  - perform data science tasks

- `import pandas as pd`: importing pandas

## data structures

- Series: single column from a Dataframe
  - `s = pd.Series([10, 20, 30])`: create a Series
- DataFrame: main data structure in pandas

```python
data = {
    'Name': ['Ana', 'Bruno', 'Carlos'],
    'Age': [23, 35, 31],
    'City': ['SP', 'RJ', 'MG']
}
df = pd.DataFrame(data)
print(df)
```

- read from file: `df = pd.read_csv('data.csv')`
- save to csv: `df.to_csv('output.csv', index=False)`

## basic commands

- `df.head(n)`: show the first `n` rows (`n` is optional)
- `df.tail(n)`: show the last `n` rows (`n` is optional)
- `df.shape`: (rows, columns)
- `df.info()`: general info
- `df.describe()`: show statistics


- `df['name']`: select one column
- `df[['name', 'age']]`: print multiple column

- `df.iloc[0]`: select row based on index
  - `df.iloc[1][0]`: row 2, column 1
  - `df.iloc[1, 0:1]`: slices
- `df.loc[0]`: select row/column (label-based)
  - `df.loc[0, 'Name']`: row 0, column `Name`

- `df[df['age'] > 30]`: filter by condition

## operations

- `df['age'].mean()`: mean
- `df['city'].value_counts()`: counts how many times each unique value appears in a Series
  - returns a Series sorted in descending order by default
- `df.sort_values('age')`: sort by age

## cleaning data

- `df.isnull()`: checks for missing null values in your dataframe
- `df.dropna()`: remove lines with null values
- `df.fillna(0)`: fill null values with `0`

