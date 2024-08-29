# DeMuddler Project

Welcome to the DeMuddler Project! This tool is designed as a companion to the existing [Muddler project](https://github.com/demonnic/muddler). While Muddler compiles a collection of JSON files and a specific folder structure into a single `.mpackage` file, DeMuddler does the reverse. It takes an `.mpackage` file and extracts its contents to create a Muddler project structure.

## Installation (pre-compiled binaries)

For pre-compiled binaries, please see the [Releases](https://github.com/Edru2/DeMuddler/releases) section.


## Installation (building the project)

For those interested in building DeMuddler themselves. 

### Getting Started

Before you begin, ensure you have Go (Golang) installed on your system. If you don't have Go installed, you can download it from [the official Go website](https://golang.org/dl/).


- **Clone the Repository:** To get started, clone this repository to your local machine. You can do this by running the following command in your terminal or command prompt:
    

```sh
git clone https://github.com/Edru2/DeMuddler/
```
    
-   **Build the Project:** Navigate to the directory where you cloned the repository. Once you are in the project's root directory, compile the project by running:
    
```sh
go build .
```
    
-    This command will compile the source code into an executable file, which you will find in the same directory.
    

## Usage

To use DeMuddler, simply use the command line interface. Here's the basic usage:

```sh

de-muddler -f [filename]
``` 

Replace `[filename]` with the name of your `.mpackage` file. For example, if your file is named `example.mpackage`, you would use:

```sh

de-muddler -f example.mpackage
``` 

This command will process the specified file using DeMuddler.

Looking to use Github and CI to publish your package? Have a look at [this guide](https://mud.gesslar.dev/muddler.html).

## Contributing and Issues

This early version of DeMuddler might have bugs or issues. 
Your feedback is crucial. 
If you notice anything or have improvement ideas, please share them in the [Issues](https://github.com/Edru2/DeMuddler/issues) section. 
All input is welcome!

## Additional Notes for Beginners

-   **What is Go (Golang)?** Go, also known as Golang, is an open-source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike.
    
-   **What are `.mpackage` files?** `.mpackage` files are specialized packages used within the Mudlet client, a popular platform for MUD (Multi-User Dungeon) gaming. These files can be created directly in Mudlet or through the Muddler tool. Normally, an `.mpackage` file is a bundle containing scripts, assets, and other resources crucial for enhancing the Mudlet gaming experience. DeMuddler is tailored to interact with these `.mpackage` files, enabling users to extract and reconstruct a project structure that can be further used with Muddler.
    
-   **Need Help?** If you have questions about using DeMuddler or need assistance with anything related to Muddler and DeMuddler, the Mudlet community is a great resource. Known for its welcoming and supportive nature, the community can provide valuable insights and help. For specific support and to engage with fellow users, join the Mudlet Discord server at [Mudlet Discord](https://discordapp.com/invite/kuYvMQ9).

