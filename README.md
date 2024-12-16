# ASCII Art Web

## Description

**Ascii-art-web** is a web application that converts plain text into ASCII art banners. Users can choose from three available banners: `standard`, `shadow`, and `thinkertoy`. The result is displayed on a new webpage. Additionally, the application provides two extra features:

- Users can save the current contents of the textarea (input text) to their device.
- Users can download a text file containing the current ASCII art output.

## Authors


- Simo ([@maidi](https://learn.zone01oujda.ma/git/maidi/ascii-art-web)) (leader)
- Abderrahmane Ben Yahya ( @abenyahya )

---

## Usage

To run the server, use:

```bash
go run . -ws
```

> **Note:** Without the `-ws` flag, the CSS file will not be supported. `-ws` stands for "with style."

If you omit the flag, run the following:

```bash
go run .
```

---

## Features

### Main Page
- A **textarea** for inputting the text to convert.
- Buttons:
  - **Submit**: Sends the text for conversion and displays the result on a new page.
  - **Clear**: Clears the contents of the textarea.
  - **Save**: Saves the current textarea contents to the user's device using localStorage.

### Result Page
- Displays the converted ASCII art.
- Buttons/Links:
  - **Back**: Returns to the main page.
  - **Download**: Downloads the ASCII art result as a `.txt` file.

### Error Handling
- **404 Not Found:** Displays a custom page if the user accesses a nonexistent route.
- **500 Internal Server Error:** Displays a page when a server-side issue occurs.
- **400 Bad Request:** If a user sends:
  - An empty textarea.
  - A request with an invalid or nonexistent banner name.
  - A direct request to the result page without proper form submission.
- Display an appropriate message on page in cases related to status 400 ...  

---

## Algorithm

1. **Validation:**
   - A function checks if the input text corresponds to a valid banner.

2. **File Handling:**
   - Loads ASCII art banner files (`standard`, `shadow`, or `thinkertoy`) into memory as maps for efficient processing.

3. **Conversion:**
   - Each line of the input text is passed through a processing function.
   - The function assembles the ASCII art line by line, starting from the top of each character and adding rows until the bottom of the character.

4. **Result Rendering:**
   - The server generates the ASCII art and sends it back to be displayed on the result page.

5. **Additional Features:**
   - Textarea content can be saved locally using `localStorage`.
   - The result can be downloaded as a `.txt` file.

---

## Steps to Use

1. Open the application by navigating to:
   ```
   http://localhost:8081/
   ```
   (display normal page if the backend server and all required files are running correctly.)

2. Enter text in the **textarea** field and select a banner style from the available options.

3. Click **Submit** to generate the ASCII art and view it on the result page.

4. Use the buttons as needed:
   - **Save:** Saves the textarea content locally.
   - **Clean text:** Clears the textarea.
   - **Back to form:** Returns to the main page.
   - **Download:** Downloads the ASCII art as a `.txt` file.

5. If you attempt to access nonexistent pages or routes, the application will display appropriate error messages:
   - **404 Not Found** for invalid routes.
   - **500 Internal Server Error** for server issues.
   - **400 Bad Request** for invalid or improperly sent input data.

---

## Example Error Scenarios 

1. Submitting an empty textarea results in a `400 Bad Request` error, visible in developer tools.
2. Sending an invalid banner name also triggers a `400 Bad Request` response.
3. Navigating directly to the result page URL without proper submission returns a `400 Bad Request` response.
4. Accessing an invalid route, such as `http://localhost:8081/unknown`, shows a `404 Not Found` page.
5. case error of server in page => `500 server error...` (but this is out frame of user).

