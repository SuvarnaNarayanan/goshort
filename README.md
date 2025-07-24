# goshort
Goshort is a URL link shortener with advanced features like click statistics while being GDPR compliant. 


I am writing a URL shortening service that lets me to shorten and track click statistics. Can you write me an openapi 3.0 spec so I could autogenerate server stubs from the api specification? 

# Requirements 

I'm thinking of using some sort of a short code that follows a general url like this:

Users sign in with google OAuth to use the service. So add all the URLs required for the OAuth authentication. 

GET {domain}/tylks --> this should redirect the browser to the original url.
GET {domain}/tylks/info --> Shows statistics like when it was clicked and everything. 

POST {domain}/tylks/report --> report the URL - requires a field for the user to write down their explanation.
POST {domain}/short --> shorten the URL and return the shortened url. 

GET {domain}/list/short --> list all the shortened URLs vs original URLs. 
GET {domain}/list/campaign --> All the campaigns. 

--- To be used programmatically by a power user ---

All the API's mentioned above can also be used with an API key created by the user. 

so {domain}/key/* 

where * can be any of the above url paths. The key must be included in the header. 

# Approach

**Initial Problem:** You needed a strategy for your URL shortening service to avoid running out of unique 5-character short codes without having to delete old URLs.

**Key Points & Proposed Solutions:**

1.  **Expanding Capacity:** We established that a 5-character code using 64 URL-safe characters (alphanumeric (a-z, A-Z, 0-9) +  '-' and '_') yields over 1 billion combinations. The primary solutions to avoid exhaustion are:
    *   **Increasing Code Length:** Moving to 6 characters provides a massive increase to ~68 billion combinations.
    *   **Dynamic Length:** The recommended approach is to start with 5-character codes and automatically switch to generating 6-character codes when the 5-character space is nearly full.

2.  **Detecting Exhaustion:** To know when the space is nearly full, you should:
    *   Monitor the percentage of used codes against the total possible.
    *   Set a threshold (e.g., 80% full) to trigger the switch to longer codes.

3.  **Efficient Counting:** To avoid slow database queries (`COUNT(*)`), the most efficient solution is to maintain a **dedicated counter** in a fast key-value store like Redis or a separate database table, which is incremented with each new URL.

4.  **Efficient Collision Handling:** To manage the increasing number of collisions as the space fills up:
    *   Rely on a **database unique constraint** on the short code column for the fastest possible collision detection.
    *   Implement a simple **retry mechanism** in your application: if an insert fails due to a collision, generate a new code and try again.
    *   Proactively switching to a longer code length at your defined threshold is the best way to keep collision probability low and performance high.
                              
            
                              
            
        