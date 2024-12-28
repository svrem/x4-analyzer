# X4 Analyzer

The X4 Analyzer is a tool that allows you to analyze that data contained in a [X4: Foundations](https://www.egosoft.com/games/x4/info_en.php) save file. The program consists of two parts: a Python script that reads the save file and stores the useful data in a SQLite database and a web interface written in Go that reads, parses and displays the extracted data.

The frontend is written in HTML and Javascript and uses [tailwindcss](https://tailwindcss.com/) and some components from [Flowbite](https://flowbite.com/). The navigation is done using [htmx](https://htmx.org/). 

## Usage

### Requirements

- [Python 3](https://www.python.org/downloads/)
- [Go](https://golang.org/)

The Python script that reads the save file doesn't have any dependencies. The Go web interface only uses the standard library and a [SQLite3](https://github.com/mattn/go-sqlite3) package.

So, you only need to clone the repository and you are good to go:

```bash
git clone https://github.com/svrem/x4-analyzer.git && cd x4-analyzer
```

### Running the Python script

To obtain the save file you want to analyze, you need to copy it from the game's save directory. The save files are located in the
`C:\Users\USERNAME\Documents\Egosoft\X4\<user-id>\save\` directory. The save files are compressed XML files with the `.xml.gz` extension, so you need to decompress them before running the Python script. 

To extract the data from a save file, you need to run the Python script `parse.py` with the path to the save file as an argument. For example:

```bash
python parse.py /path/to/savefile.xml
```

This will create a SQLite database file named `data.db` in the current directory with the extracted data. 

### Running the web interface

To run the web interface, you need to run the following command:

```bash
go run .
```

This will start a web server on `localhost:8080`. You can access the web interface by opening a web browser and navigating to `http://localhost:8080`.

### Running with Docker

Alternatively, you can run the web interface using Docker. First, you need to build the Docker image:

```bash 
docker build -t x4-analyzer .
```

Then, you can run the Docker container:

```bash
docker run -p 8080:8080 x4-analyzer
```

This will start a web server on `localhost:8080`. You can access the web interface by opening a web browser and navigating to `http://localhost:8080`.