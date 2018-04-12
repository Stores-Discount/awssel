###### OLST / TOOLING
# AWSSEL (AWS Ssm Env Loader)


A tool that load all env vars stored in AWS SSM Parameter Store, for a given service.

See [Our Moviations doc](docs/awssel/motivations.md) for more information about why we build this tool and understand how things work under the hood.



## Contents

* [Install]
* [Usage]
* [Examples]
* [TODO]
* [Tests]
* [Guidelines]
* [Contribute]
* [FAQ]
* [References]
* [Maintainers]
* [License]

<br>

## Install

To install the tool from your shell, run:

```
go get -u github.com/Stores-Discount/awssel
```

## Usage

```
awssel <command> [options]
```

### Commands

**Awssel** has a only 2 commands : 



Command | Description | Options
---------|----------|---------
 help | Show command help text | -
 load | Load env vars for a service | See [Load options](docs/awssel/cmd/load.md)


## Examples

Get all environment variables for `proxy-web` service

```bash
awssel load --service-name proxy-web
```

## TODO

* [ ] List errors about AWS Credentials (FAQ)
* [ ] Add a section explaining how to run this tool in an AWS

## Tests

**awssel** is battle tested. We use the [localstack](https://github.com/localstack/localstack) project to mock AWS services.
In fact, before launch our tests, we prepare our test environment by starting an AWS SSM Mock; running at `http://localhost:4583`.

Preparing the test environment, involves executing the following test fixtures: 

- Start localstack AWS SSM Service
  ```bash
  rake code:test:prepare  # Start testing env
  ```
- Add test entries into the parameter store. Look into the file [tasks/populate.rb](tasks/populate.rb) for more details.
  ```bash
  rake code:test:seed  # Populate SSM with test values
  ```

At this point, the test environment is ready and you can run the following command to launch our tests:

```bash
rake code:test:run # Run all tests
```

##### CAUTION
> You choose to focus on testing the business logic behind **awssel** and not the CLI UI.
> Feel free to contribute; see [Contribute](#contribute) section to know how to contribute.


### QA


## Guidelines

See the [guidelines doc].

## Contribute

See the [contributing doc].

## FAQ

See the [faq doc].

## References

* [Golang regex syntax](https://github.com/google/re2/wiki/Syntax)

## Maintainers

* Lionel T. [@lktslionel](https://twitter.com/lktslionel)

## License
 
[MIT license]


[Changelog]: docs/CHANGELOG.md
[contributing doc]: docs/CONTRIBUTE.md
[guidelines doc]: docs/GUIDELINES.md
[faq doc]: docs/FAQ.md
[MIT license]: LICENSE
[Install]: #Install
[Usage]: #Usage
[Examples]: #Examples
[TODO]: #TODO
[Tests]: #Tests
[Guidelines]: #Guidelines
[Contribute]: #Contribute
[FAQ]: #FAQ
[References]: #References
[Maintainers]: #Maintainers
[License]: #License