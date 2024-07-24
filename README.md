# Stylize

## Project Description

Stylize is a web-based tool designed to enhance the user experience by making a website more appealing, interactive, and intuitive. This project integrates various modern web technologies and practices to ensure that the website is user-friendly, responsive, and provides comprehensive feedback to users. Stylize is an improvement of ASCII Art Web which is a web based tool developed in Go and HTML that converts text input into ASCII art using predefined character sets and displays it on the browser. This tool also allows users choose from multiple ASCII art styles to enhance the visual representation of their text.

### Objectives

Stylize focuses on providing an engaging and seamless user experience. The primary objective of this project is to create a user-friendly and visually appealing website that incorporates the following features:

- Improved interactivity and intuitiveness.
- Enhanced user-friendliness.
- Increased feedback to users.

### Key Learning Outcomes

- The basics of human-computer interaction.
- The fundamentals of CSS.
- Linking CSS and HTML.
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
    <li><a href="#optional-styles">Optional Styles</a></li>
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
- **error.go**: A package that handles error page to allow for effective communication of any errors that occur

### Built With

- Go Programming Language
- CSS
- JavaScript
- HTML

## Getting Started

To get started with Stylize, follow the instructions below.

### Prerequisites

Before running the program, ensure that you have the following prerequisites:

- Go installed on your machine.
- Basic understanding of Go programming language.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Vincent-Omondi/stylize.git
    ```

2. Navigate to the project directory:

    ```sh
    cd stylize
    ```

## Usage

To use ASCII Art Web, follow these steps:

1. Run the program with the following command:

    ```sh
    go run . 
    ```
2. Open your browser and on a new tab, go to ```http://localhost:8080 ```


 <img src="web_image.png" alt="Home page" width="600" height="400">


3. Input your text, choose your preferred style of display(banner) and then submit by pressing the generate ascii art button.

## Optional Styles

Stylize supports multiple ASCII art styles, including:

- **Standard**: A basic ASCII art style.
- **Shadow**: ASCII art with shadow effects.
- **Thinkertoy**: ASCII art with a playful design.
- **roman_space.txt**
- **roman.txt**

## Expected Output

Input Text: Hello

Instance 1. standard(default) 
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                      
```                                   

Instance 2. shadow
```
                              
_|    _|          _| _|          
_|    _|   _|_|   _| _|   _|_|   
_|_|_|_| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
                                 

```

Instance 3. thinkertoy

```
                 
o  o     o o     
|  |     | |     
O--O o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                                                                                  

```
Instance 4. roman

```
                 
ooooo   ooooo           oooo  oooo            
`888'   `888'           `888  `888            
 888     888   .ooooo.   888   888   .ooooo.  
 888ooooo888  d88' `88b  888   888  d88' `88b 
 888     888  888ooo888  888   888  888   888 
 888     888  888    .o  888   888  888   888 
o888o   o888o `Y8bod8P' o888o o888o `Y8bod8P' 
                                              
                                                                                
```
             
## Roadmap

The following features are planned for future releases:
- Add support for color.
- Add feature to specify text alignment
- Integration with third-party ASCII art libraries.
- Enhanced CSS animations and transitions.

## Contributing

Contributions to stylize are welcome! If you'd like to contribute, please follow these steps:

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
