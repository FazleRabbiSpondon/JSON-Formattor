# JSON-Formattor

So, I heard if you marshall a JSON 2 times, you end up with formats using "\" embedded into every field. Somewhat like the following
"extra": "{\"payload\":{\"id\":688},\"response\":{\"status\":true,\"message\":\"Success\",\"user\":{\"id\":688,\"name\":\"system admin\",\"mobileNumber\":\"01911223399\",\"email\":\"sysuser@gmail.com\",\"userType\":\"CMS\"}}}"
}

It renders it incompatible to beatify with the avaialable JSON beautifiers. The nested structure makes it more difficult and unreadable, needless to say inconceivable. 

Until now, I was using a ChatGPT to format it. But after a while I noticed the GPT changes the data, some of the which was important for me. After repeated attempt to convince GPT to not alter my data, rather to format it only, it didn't listen and I gave up. 

So, I write a code to ommit the "\"s and add `"` where necessary so that I can beautify the JSON with already avbailable beautifier.

Just Paste your mixed up (with \) JSON into a file named "myfile.txt" which resides in the same directory of the Go file. Run the Go file. The "myfile.txt" would be changed into a formatted proper JSON which can then be beautified with any JSON beautifier.
