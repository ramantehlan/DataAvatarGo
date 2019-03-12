# DataAvatarGo

DataAvatarGo aims to perform reincarnation of personal data entities in unstructured data sets. This project was created in March-2019, for [IIM Ahmedabad’s second hackathon](https://sites.google.com/iimahd.ernet.in/hackathon-icadabai2019/home), organized in collaboration with IBM, at [ICADABAI 2019](https://sites.google.com/iimahd.ernet.in/hackathon-icadabai2019/home).

## Problem Statement

> A famous study found that with the date of birth, gender, and five-digit zip code you can uniquely identify about 87% of Americans.

After **Data Protection Regulations** like [GDPR](https://www.google.com/url?q=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FGeneral_Data_Protection_Regulation&sa=D&sntz=1&usg=AFQjCNFZDoQ-bJV5QtpyAldQEG1GWL5l3w) in the European Union and the [Personal Data Protection Bill 2018](http://www.google.com/url?q=http%3A%2F%2Fmeity.gov.in%2Fwritereaddata%2Ffiles%2FPersonal_Data_Protection_Bill%2C2018.pdf&sa=D&sntz=1&usg=AFQjCNGdvYBwwcQChDsXip2rZzziO2afpA) in India, there is increasing regulatory backing for our privacy, and the protection of our personal data. It is also a big challenge for companies to protect the personal data entity of customers. A Personal Data Entity is any entity that can help identify or profile a person, and hence needs to be protected.  

#### Solution

A algorithm to automatically impute values for the redacted portions in a text, which are known to have contained Personal Data Entities (PDEs), then we can protect the PDEs.

#### Pseudo code

```

...

```


## Usage

Write about usage here

## Pre-Requisites

- Go Lang
- Terminal
- Text Editor

## File Structure

```
.
├── data
│   ├── all_pde_types.json
│   ├── dataset.json
│   ├── documents.json
│   ├── sliceDataset.json
│   ├── sliceSampleSubmitFile.json
│   └── types.json
├── DataAvatarGo
├── jsonLib.go
├── jsonStuct.go
├── main.go
└── README.md
```

No | File/Folder Name | Purpose |
---|------------------|---------|
1 | `submission.json` | This is the final result file, which is to be submitted
2 | `Data` | folder to store all the json files
3 | `dataset.json` | Data set provided
4 | `document.json` | Original file from which the dataset is created
5 | `types.json` | Few selected PDEs types to recognize
6 | `all_pde_types.json` | All the provided PDEs types
7 | `README.md` | Current file you are reading
8 | `main.go` | Main file for the program
9 | `jsonStruct.go` | It holds the struct used for json reading and writing
10 | `jsonLib.go` | It holds all the functions used for json parsing and writing
11 | `sliceDataset.json` | Just for testing
12 | `sliceSampleSubmitFile.json` | Just for testing

## Why Go Language?

There are many reasons to use `Go Lang` as a language of choice, some of the important ones are mentioned below.

#### Concurrency

Go is built with keeping concurrency in mind. Go has goroutines instead of threads which is not just fast in starting up, but use less memory.

#### Binary Code

Go is compiled and not interpreted, so it runs directly on underlying hardware.

#### Maintainable - Reliable - Scalable

Go is one of the easiest language to learn, with simple and clean syntax. Though, it intentionally leaves many features out, Like:

- No Classes
- No Inheritance
- No constructors
- No annotations
- No generics
- No exceptions

You may find many of the features crucial, but you can still code your application in go, with few extra lines. That will not just make you code more clean, but more easy to understand. Go is a statically typed and compiled, which makes it reliable and scalable.


## Resources

- [jdkato/prose](https://github.com/jdkato/prose)
- [Prodigy prose radically efficient machine teaching in go](https://medium.com/@jdkato/prodigy-prose-radically-efficient-machine-teaching-in-go-93389bf2d772)
- [Workflow named entity recognition](https://prodi.gy/docs/workflow-named-entity-recognition)
- [Why should you learn go](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)


## License

Not yet decided
