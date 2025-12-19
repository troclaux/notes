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

new_rows = [
    {"name": "Diana", "age": 28}
]

df = pd.concat([df, pd.DataFrame(new_rows)], ignore_index=True)

# name  age city
# 0     Ana   25   SP
# 1   Bruno   30   RJ
# 2  Carlos   35   BH
# 3   Diana   28  NaN

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

## Pandas basics

- Create data: from dict/list/CSV

```python
import pandas as pd
df = pd.DataFrame({"a": [1, 2], "b": ["x", "y"]})
s = pd.Series([10, 20, 30], name="values")
df = pd.read_csv("file.csv")  # read_json/read_parquet/read_excel similar
```

- Inspect: `df.head()`, `df.tail()`, `df.info()`, `df.describe()`, `df.shape`, `df.dtypes`, `df.columns`, `df.index`
- Select columns/rows: `df["a"]`, `df[["a", "b"]]`, `df.loc[row_labels, col_labels]`, `df.iloc[row_idx, col_idx]`, boolean mask `df[df["a"] > 1]`, query `df.query("a > 1 and b == 'x'")`
- Assign/mutate: `df["c"] = df["a"] * 2`, `df.assign(c=lambda d: d["a"] * 2, d=5)`, `df.rename(columns={"a": "alpha"})`, `df.drop(columns=["b"])`
- Filtering/missing: `df.dropna()`, `df.fillna(0)`, `df.isna().sum()`, `df["a"].between(1, 5)`, `df["b"].str.contains("x")`
- Sorting: `df.sort_values(["a", "b"], ascending=[True, False])`, `df.sort_index()`
- Grouping/aggregation:

```python
df.groupby("b")["a"].agg(["mean", "sum", "count"])
df.groupby(["b"]).agg(total=("a", "sum"), avg=("a", "mean"))
```

- Pivot/reshape: `df.pivot(index="row", columns="col", values="val")`, `df.pivot_table(values="val", index="row", columns="col", aggfunc="mean")`, `df.melt(id_vars=["id"], var_name="metric", value_name="value")`
- Join/merge/concat:

```python
pd.concat([df1, df2], axis=0, ignore_index=True)  # stack rows
pd.concat([df1, df2], axis=1)  # align on index
df1.merge(df2, on="key", how="inner")  # left/right/outer available
```

- Apply/transform: `df["a"].apply(lambda x: x + 1)`, `df.applymap(func)` (elementwise), `df.apply(func, axis=1)` (rowwise), `df.transform({"a": ["mean", "std"]})` for group-wise broadcasts
- String/datetime: `df["text"].str.lower()`, `.str.contains("pattern")`; `df["date"] = pd.to_datetime(df["date_str"])`, accessors like `df["date"].dt.year`; resample: `df.resample("M").sum()` with datetime index
- Indexing tricks: set/reset index `df.set_index("id")`, `df.reset_index()`, multi-index from columns: `df.set_index(["a", "b"])`
- Window ops: rolling/expanding: `df["a"].rolling(3).mean()`, `df["a"].expanding().sum()`
- Describe types: numeric vs categorical; convert: `df["col"] = df["col"].astype("category")`, `pd.to_numeric(..., errors="coerce")`
- Export: `df.to_csv("out.csv", index=False)`, `.to_parquet`, `.to_json`, `.to_excel`
- Performance tips: prefer vectorized ops over Python loops; use `.loc` for assignment to avoid SettingWithCopy warnings; use categorical for low-cardinality strings; leverage `read_csv` `dtype` and `usecols` to speed up loads

Quick example end-to-end:

```python
df = pd.read_csv("sales.csv")
df["revenue"] = df["units"] * df["price"]
top = (
    df[df["revenue"] > 1000]
    .groupby("region")
    .agg(total_revenue=("revenue", "sum"), orders=("id", "count"))
    .reset_index()
    .sort_values("total_revenue", ascending=False)
)
top.to_csv("top_regions.csv", index=False)
```
