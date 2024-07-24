# Stylize

## Project Description

ASCII Art Web is a web based tool developed in Go and HTML that converts text input into ASCII art using predefined character sets and displays it on the browser. This tool also allows users choose from multiple ASCII art styles to enhance the visual representation of their text.

### Objectives

The primary objective of this project is to provide a flexible and user-friendly tool for generating ASCII art from text input. It offers the following features:

- Conversion of text input into ASCII art using customizable character sets.
- Multiple ASCII art styles to choose from, including standard, shadow, and thinkertoy.


<!-- TABLE OF CONTENTS -->
<details>
  <summary style="font-weight: bold; font-size: 1.4em;" >Table of Contents</summary>
  <ol>
    <li>
      <a href="#implementation-details">Implementation details</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#optional-ascii-art-styles">Optional ASCII Art Styles</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#authors">Authors</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## Implementation details

ASCII Art Web is designed to provide users with a simple yet powerful tool for creating ASCII art from text input. It utilizes a set of ASCII characters to represent text in a visually appealing manner. The project consists of three main components:

- **main.go**: The entry point of the program, responsible for creating local server on port 8080 and routing URL paths.
- **loadascii.go**: A package that handles the loading of ASCII characters from file.
- **printascii.go**: A package that prints the ASCII art to the specified output file.
- **asciiart.go**: A package that handles form action, loading ascii characters and generating ascii art
- **index.go**: A package that handles the landing page of the web tool

### Built With

- Go Programming Language
- HTML

## Getting Started

To get started with ASCII Art Web, follow the instructions below.

### Prerequisites

Before running the program, ensure that you have the following prerequisites:

- Go installed on your machine.
- Basic understanding of Go programming language.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Vincent-Omondi/ascii-art-web.git
    ```

2. Navigate to the project directory:

    ```sh
    cd ascii-art-web
    ```

## Usage

To use ASCII Art Web, follow these steps:

1. Run the program with the following command:

    ```sh
    go run . 
    ```
2. Open your browser and on a new tab, go to ```http://localhost:8080 ```
3. Input your text, choose your preferred style of display(banner) and then submit by pressing the generate ascii art button

## Optional ASCII Art Styles

ASCII Art Output supports multiple ASCII art styles, including:

- **Standard**: A basic ASCII art style.
- **Shadow**: ASCII art with shadow effects.
- **Thinkertoy**: ASCII art with a playful design.

## Expected Output

Input Text: Hello

Instance 1. shadow 
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                      
```                                   

Instance 2. thinktertoy
```
                              
_|    _|          _| _|          
_|    _|   _|_|   _| _|   _|_|   
_|_|_|_| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
                                 

```

Instance 3. standard

```
                 
o  o     o o     
|  |     | |     
O--O o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                                                                                  

```
## Roadmap

The following features are planned for future releases:
- Add support for color.
- Add feature to specify text alignment
- Integration with third-party ASCII art libraries.

## Contributing

Contributions to ASCII Art Output are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your fork.
5. Open a pull request to merge your changes into the main branch.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Authors

- **[X - @vinomondi_1](https://x.com/vinomondi_1)**
- **[X - @oumaphilip01](https://x.com/oumaphilip01)**
- **[Github - Vincent](https://github.com/Vincent-Omondi/)**
- **[Github - Philip38](https://github.com/Philip38-hub)**

<p align="right">(<a href="#ascii-art-web">back to top</a>)</p>


## Acknowledgments

Special thanks to the creators of the ASCII art character sets used in this project.




Project Description

Stylize is a web-based tool designed to enhance the user experience by making a website more appealing, interactive, and intuitive. This project integrates various modern web technologies and practices to ensure that the website is user-friendly, responsive, and provides comprehensive feedback to users.
Objectives

The primary objective of this project is to create a user-friendly and visually appealing website that incorporates the following features:

    Improved interactivity and intuitiveness.
    Enhanced user-friendliness.
    Increased feedback to users.

Key Learning Outcomes

This project helps in understanding:

    The basics of human-computer interaction.
    The fundamentals of CSS.
    Linking CSS and HTML.

Allowed Packages

    Only standard Go packages are allowed.


Stylize
Project Description

Stylize is a web-based tool designed to enhance the user experience by making a website more appealing, interactive, and intuitive. This project integrates various modern web technologies and practices to ensure that the website is user-friendly, responsive, and provides comprehensive feedback to users.
Objectives

The primary objective of this project is to create a user-friendly and visually appealing website that incorporates the following features:

    Improved interactivity and intuitiveness.
    Enhanced user-friendliness.
    Increased feedback to users.

Key Learning Outcomes

This project helps in understanding:

    The basics of human-computer interaction.
    The fundamentals of CSS.
    Linking CSS and HTML.

Allowed Packages

    Only standard Go packages are allowed.

<!-- TABLE OF CONTENTS -->
<details>
  <summary style="font-weight: bold; font-size: 1.4em;">Table of Contents</summary>
  <ol>
    <li>
      <a href="#implementation-details">Implementation Details</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#authors">Authors</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>
Implementation Details

Stylize focuses on providing an engaging and seamless user experience. The project architecture consists of several components that ensure the smooth functioning of the website.
Built With

    Go Programming Language
    Node.js
    Express.js
    Redis
    CSS
    HTML

Getting Started

To get started with Stylize, follow the instructions below.
Prerequisites

Ensure that you have the following prerequisites before running the program:

    Go installed on your machine.
    Node.js and npm installed on your machine.
    Redis installed and running on your machine.
    Basic understanding of Go and JavaScript programming languages.

Installation

    Clone the repository:

    sh

git clone https://github.com/yourusername/stylize.git

Navigate to the project directory:

sh

cd stylize

Install the necessary Node.js packages:

sh

npm install

Configure the Redis server details in default.json:

json

    {
      "redis": {
        "host": "127.0.0.1",
        "port": 6379
      }
    }

Usage

To use Stylize, follow these steps:

    Start the server with the following command:

    sh

node index.js

Open your browser and navigate to:

sh

    http://localhost:3000

    Interact with the various features on the website, ensuring a consistent, responsive, and interactive user experience.

Roadmap

The following features are planned for future releases:

    Enhanced CSS animations and transitions.
    Additional interactive elements such as modals and tooltips.
    Integration with external APIs for dynamic content.

Contributing

Contributions to Stylize are welcome! If you'd like to contribute, please follow these steps:

    Fork the repository.
    Create a new branch for your feature or bug fix.
    Make your changes and commit them with descriptive messages.
    Push your changes to your fork.
    Open a pull request to merge your changes into the main branch.

License

This project is licensed under the MIT License. See the LICENSE file for more details.
Authors

    GitHub - YourUsername

<p align="right">(<a href="#stylize">back to top</a>)</p>
Acknowledgments

Special thanks to the creators of the various tools and libraries used in this project, including Node.js, Express.js, and Redis.
File Descriptions
Configuration Files

    default.json: Contains default configuration settings for the application, including Redis server details​​.
    configure.js: Merges default and custom configurations using mixme​​​​.

Database

    dbClient.js: Establishes a connection to the Redis database and exports the Redis client instance for use throughout the application​​.

Server and Routing

    index.js: Main entry point for the Node.js server. Sets up the Express.js application, connects to Redis, and defines routes​​.
    user.controller.js: Contains logic for user-related operations such as creating and retrieving users​​.
    user.router.js: Defines REST API routes for user operations and includes test cases for these routes​​.

Frontend

    index.html: The main HTML file for the website, containing the structure and content of the homepage. Includes CSS for styling the webpage​​.

Package Configuration

    package-lock.json: Specifies the versions of the installed Node.js packages and their dependencies​​.