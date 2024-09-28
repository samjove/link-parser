## HTML Link Parser

This program parses links and their labels from a given block of HTML code. 

It can be run by cloning the repo, building with `go build` and running the resulting `link-parser.exe`.

#### Example input (included in main.go):

    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Sample HTML with Links</title>
    </head>
    <body>
        <h1>Welcome to My Website</h1>
        <p>This is a sample paragraph with a link to <a href="https://www.example.com">Example</a>.</p>
        <p>Check out my <a href="https://www.portfolio.com">Portfolio</a> for more projects.</p>
        <p>For more information, visit <a href="https://www.documentation.com">Documentation</a>.</p>
        <footer>
            <p>Contact us at <a href="mailto:info@example.com">info@example.com</a>.</p>
            <p>Follow us on <a href="https://www.socialmedia.com">Social Media</a>.</p>
        </footer>
    </body>
    </html>

#### Example output:
    [
        {URL:https://www.example.com Label:Example} 
        {URL:https://www.portfolio.com Label:Portfolio} 
        {URL:https://www.documentation.com Label:Documentation} 
        {URL:mailto:info@example.com Label:info@example.com} 
        {URL:https://www.socialmedia.com Label:Social Media}
    ]