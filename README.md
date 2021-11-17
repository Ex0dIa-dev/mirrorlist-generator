### mirrorlist-generator

****

**Build**

```bash
go build main.go
```

****

**Usage**

```bash
#interactive mode
./mirrorlist-generator -i

#one command mode
./mirrorlist-generator -countries="" -protocols="" -ipversions=""

#listing data
./mirrorlist-generator -lc	#printing all available countries
./mirrorlist-generator -lp 	#printing all available protocols
./mirrorlist-generator -li 	#printing all available ipversions
```

**Example**

```bash
./mirrorlist-generator -countries="all" -protocols="http" -ipversions="4"
./mirrorlist-generator -countries="IT" -protocols="http" -ipversions="4"

#you can enter multiple countries, etc..
./mirrorlist-generator -countries="IT,EN" -protocols="http,https" -ipversion="4,6"
```

****

**Flags:**

- **-i** ==> *interactive mode*
- **-countries** ==> *enter here countries (ordered and comma-separated)*
- **-protocols** ==> *enter here protocols (ordered and comma-separated)*
- **-ipversions** ==> *enter here ipversions (ordered and comma-separated)*
- **-lc** ==> *print all available countries*
- **-lp** ==> *print all available protcols*
- **-li** ==> *print all available ipversions*

