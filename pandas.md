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
- Dataframe: main data structure in pandas

```python
data = {
    'Nome': ['Ana', 'Bruno', 'Carlos'],
    'Idade': [23, 35, 31],
    'Cidade': ['SP', 'RJ', 'MG']
}
df = pd.DataFrame(data)
print(df)
```

- read from file: `df = pd.read_csv('dados.csv')`
- save to csv: `df.to_csv('saida.csv', index=False)`

## basic commands

- `df.head()`: show first columns
- `df.tail()`: show last columns
- `df.shape`: (rows, columns)
- `df.info()`: general info
- `df.describe()`: show statistics


- `df['name']`: print one column
- `df[['name', 'age']]`: print multiple column

- `df.iloc[0]`: select row based on index
  - `df.iloc[1, 0]`: row 2, column 1
- `df.loc[0]`: select row/column (label-based)
  - `df.loc[0, 'Name']`: row 0, column `Name`

- `df[df['age'] > 30]`: filter by condition

## operations

- `df['age'].mean()`: mean
- `df['city'].value_counts()`: counts how many times each unique value appears in a Series
- `df.sort_values('age')`: sort by age

## cleaning data

- `df.isnull()`: checks for missing null values in your dataframe
- `df.dropna()`: remove lines with null values
- `df.fillna(0)`: fill null values with `0`

