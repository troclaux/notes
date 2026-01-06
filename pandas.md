[python](./python.md)

# Pandas

> python library for data manipulation and analysis

- features
  - load and clean data
  - explore, filter, sort and group data
  - perform data science tasks

- `import pandas as pd`: importing pandas

## data structures

- series: single column from a Dataframe
  - `s = pd.Series([10, 20, 30])`: create a Series
- dataFrame: main data structure in pandas

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

print(df)

# name  age city
# 0     Ana   25   SP
# 1   Bruno   30   RJ
# 2  Carlos   35   BH
# 3   Diana   28  NaN
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

```python
import pandas as pd
df = pd.DataFrame({"a": [1, 2], "b": ["x", "y"]})
s = pd.Series([10, 20, 30], name="values")
df = pd.read_csv("file.csv")  # read_json/read_parquet/read_excel similar
```

- create data: from dict/list/CSV
- inspect: `df.head()`, `df.tail()`, `df.info()`, `df.describe()`, `df.shape`, `df.dtypes`, `df.columns`, `df.index`
- select columns/rows: `df["a"]`, `df[["a", "b"]]`, `df.loc[row_labels, col_labels]`, `df.iloc[row_idx, col_idx]`, boolean mask `df[df["a"] > 1]`, query `df.query("a > 1 and b == 'x'")`
- assign/mutate: `df["c"] = df["a"] * 2`, `df.assign(c=lambda d: d["a"] * 2, d=5)`, `df.rename(columns={"a": "alpha"})`, `df.drop(columns=["b"])`
- filtering/missing: `df.dropna()`, `df.fillna(0)`, `df.isna().sum()`, `df["a"].between(1, 5)`, `df["b"].str.contains("x")`
- sorting: `df.sort_values(["a", "b"], ascending=[True, False])`, `df.sort_index()`
- grouping/aggregation:

```python
df.groupby("b")["a"].agg(["mean", "sum", "count"])
df.groupby(["b"]).agg(total=("a", "sum"), avg=("a", "mean"))
```

- pivot/reshape: `df.pivot(index="row", columns="col", values="val")`, `df.pivot_table(values="val", index="row", columns="col", aggfunc="mean")`, `df.melt(id_vars=["id"], var_name="metric", value_name="value")`
- join/merge/concat:

```python
pd.concat([df1, df2], axis=0, ignore_index=True)  # stack rows
pd.concat([df1, df2], axis=1)  # align on index
df1.merge(df2, on="key", how="inner")  # left/right/outer available
```

- apply/transform: `df["a"].apply(lambda x: x + 1)`, `df.applymap(func)` (elementwise), `df.apply(func, axis=1)` (rowwise), `df.transform({"a": ["mean", "std"]})` for group-wise broadcasts
- string/datetime: `df["text"].str.lower()`, `.str.contains("pattern")`; `df["date"] = pd.to_datetime(df["date_str"])`, accessors like `df["date"].dt.year`; resample: `df.resample("M").sum()` with datetime index
- indexing tricks: set/reset index `df.set_index("id")`, `df.reset_index()`, multi-index from columns: `df.set_index(["a", "b"])`
- window ops: rolling/expanding: `df["a"].rolling(3).mean()`, `df["a"].expanding().sum()`
- describe types: numeric vs categorical; convert: `df["col"] = df["col"].astype("category")`, `pd.to_numeric(..., errors="coerce")`
- export: `df.to_csv("out.csv", index=False)`, `.to_parquet`, `.to_json`, `.to_excel`
- performance tips: prefer vectorized ops over Python loops; use `.loc` for assignment to avoid SettingWithCopy warnings; use categorical for low-cardinality strings; leverage `read_csv` `dtype` and `usecols` to speed up loads

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

## string accessor (.str)

- `.str` is the pandas "string accessor." It exposes vectorized string methods on a Series (or Index) so you can apply string logic element-wise without looping.

Quick rules

- works on Series/Index with string/object dtype (or mixed with strings)
- handles missing values; many methods accept `na=`
- most methods mirror Python string methods or regex tools

Common patterns

```python
s = df["SECURITIES"]

# contains / match
s.str.contains("AUAU", na=False)
s.str.match(r"^AU")          # regex match from start
s.str.fullmatch(r"AU\\d+")

# case / normalization
s.str.lower()
s.str.upper()
s.str.strip()                # trim whitespace

# replace / extract
s.str.replace(r"\\s+", " ", regex=True)
s.str.extract(r"(AU\\d+)", expand=False)

# split / join
s.str.split("-", n=1, expand=True)
s.str.join(",")

# length / slicing
s.str.len()
s.str[:4]
```

Filtering example

```python
df[df["SECURITIES"].str.contains("AUAU", na=False)]
```

If you are working with non-strings, convert first:

```python
df["SECURITIES"].astype("string").str.contains("AUAU", na=False)
```

## operations

- `df['age'].mean()`: mean
- `df['city'].value_counts()`: counts how many times each unique value appears in a Series
  - returns a Series sorted in descending order by default
- `df.sort_values('age')`: sort by age

## cleaning data

- `df.isnull()`: checks for missing null values in your dataframe
- `df.dropna()`: remove lines with null values
- `df.fillna(0)`: fill null values with `0`

## loc vs iloc

- `loc` is label-based
  - slices are end-inclusive
  - errors if a label is missing
- `iloc` is position-based (0-based)
  - slices are end-exclusive
  - errors if an index is out of range

```python
df = pd.DataFrame({"a": [10, 20, 30], "b": ["x", "y", "z"]}, index=["r1", "r2", "r3"])

df.loc["r2"]                      # row with label r2
df.loc["r1":"r2", ["a"]]         # label slice rows r1..r2, column a

df.iloc[0]                        # first row by position
df.iloc[1:3, 0]                   # rows 1-2, first column (end-exclusive)
df.iloc[[0, 2], [1]]              # rows 0 and 2, column at position 1 (b)
```

## conditionals

- build boolean masks to filter, combine with `&` and `|` (wrap each condition in parentheses)

```python
mask = df["a"] > 15
df[mask]
df[(df["a"] > 15) & (df["b"] == "y")]
df[df["a"].between(15, 25)]
df[df["b"].str.contains("y")]
df.query("a > 15 and b == 'y'")
df.loc[df["a"] > 15, "flag"] = True   # assign based on condition
```

## isin

- Membership test that returns a boolean mask; useful for inclusion or exclusion filters. Matches are exact (case-sensitive for strings); NaN compares False.

```python
df[df["country"].isin(["BR", "AR"])]              # keep rows where country is BR or AR
df[~df["country"].isin(["US", "CA"])]             # exclude US and CA

# DataFrame-wide: supply per-column allowed values, then collapse per row
allowed = {"status": ["new", "open"], "priority": [1, 2]}
df[df.isin(allowed).any(axis=1)]                  # any match per row
df[df.isin(allowed).all(axis=1)]                  # all listed columns match

ids = {101, 202, 303}
df[df["id"].isin(ids)]                            # set is fine for fast lookups
```
