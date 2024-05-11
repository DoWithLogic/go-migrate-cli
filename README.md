# go-migration-cli

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**go-migration-cli** is a custom command-line interface (CLI) tool built using [golang-migrate](https://github.com/golang-migrate/migrate) for managing database migrations in Go projects. This CLI simplifies the process of applying and managing database schema changes, making it easier to version-control and collaborate on database changes within your projects.

The goal of this repository is to provide a seamless solution for incorporating database migrations into your CI/CD pipelines without the need to install additional libraries or programming languages. By utilizing **go-migration-cli**, you can automate the database migration process within your pipelines, ensuring consistency and reliability across your development and deployment workflows.


## Features

- **Simplified Migration Management**: Manage database migrations effortlessly using a familiar command-line interface.
- **Support for Multiple Database Engines**: Utilize golang-migrate's support for various database engines including MySQL, PostgreSQL, SQLite, and more. [currently just support MySQL]
- **Version Control Integration**: Seamlessly integrate database migrations with your version control system (VCS) for collaborative development.
- **Customizable Workflow**: Tailor the migration workflow to fit your project's specific requirements with configurable options.
- **Easily Extensible**: Extend the CLI's functionality by integrating additional plugins or custom scripts as needed.

## Getting Started

To start using **go-migration-cli**, follow these steps:
1. build go-migrate-cli
```bash
    make build-cli
```

2. update permissions of executable file
``` bash
    sudo chmod +x mcli
```

3. move into your local bin
``` bash
    sudo cp mcli /usr/local/bin/
```

4. run migration-up
``` bash
    mcli migrate up -H 127.0.0.1:3006 -N database_name -P pwd -U root --path /database/migrations
```

5. run migration-down
``` bash
    mcli migrate down -H 127.0.0.1:3006 -N database_name -P pwd -U root --path /database/migrations
```

## Contributing

Contributions are welcome! If you encounter any issues, have feature requests, or would like to contribute enhancements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
