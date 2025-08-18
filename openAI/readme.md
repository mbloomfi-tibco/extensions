Flogo ChatGPT Client Exension
=============================

Please do not use this for your own projects at this stage as it is still in very early stage of development.  If you are a TIBCO customer and would like to use the extension or have feeback then please reach out to the repository owner or your account team. 

Overview
--------
The Flogo ChatGPT extesion acts as a client for the APIs exposed from the ChatGPT developer platform.  It is in the very early days of development and there is a focus on supporting text generation, image generation, embedings creation and storage using the Responses API. 

Mini Roadmap 
------------


| Client Activity                     | Status                        | OpenAI Go API Library Status  |
| ------------------------------------| ------------------------------| ------------------------------|
| Responses API                       | In Dev for text and Images    | v1.12.0                       |
| Images API                          | In Dev for text and Images    | v1.12.0                       |
| Embedings API                       | In Dev for text and Images    | v1.12.0                       | 
| Chat API                            | Currently out of scope        | 
| Completions API                     | Currently out of scope        |
| Realtime API                        | Currently out of scope        |  
| Assistants API                      | Currently out of scope        |
| Batch API                           | Currently out of scope        |
| Containers API                      | Currently out of scope        |
| Files API                           | Currently out of scope        |
| Fine Tunning API                    | Currently out of scope        |
| Graders API                         | Currently out of scope        |
| Moderrations API                    | Currently out of scope        | 

ChatGPT Use Cases 
-----------------


| Client Activity                     | Status                        |
| ----------- ------------------------| ------------------------------|
| Text Generation                     | In Dev                        | 
| Image Generation                    | In Dev                        |
| Audio Generation                    | Currently out of scope        |
| Deep Research                       | Currently out of scope        |  
| Embeddings                          | In Dev                        |
| Moderation                          | Currently out of scope        |

Flogo Example Use Caes 
----------------------

These are some of the example use cases we are being reviewed as part of the client extension development.

Use Flogo to get Structured Data from a model and store it in Mongo DB.

Use Flogo to generate Text and Images and then forward and send to a message broker.  

Use Flogo to Analyse images that are part of an event stream

Use flogo to first expose APIs with MCP and then use Remote MCP with an LLM.

Use flogo to do a file search

Use flogo to do a web search


OpenAI api Go Library
---------------------

To interact with the open AI API the official go library is used.  More information on the API can be found here:

https://github.com/openai/openai-go/blob/main/README.md


More Information
----------------

For more information please reach out to the repository owner.