<h1>go-reloaded test file</h1>

This test file will check your go-reloaded project against the test cases on the intra audit web page.

The test file needs to be copied to your project folder, the same folder that your main project file is located.

If your main project file is of package "main" and your project is processed in the "main()" function, then you will need to change the name of that function to GoReloaded(), or another name of your choosing (has to start with a capital letter, to make it public), and then create another function called main() where you call the original functions new name. The function names should now look like this:

![expected_func_names](https://user-images.githubusercontent.com/81628708/137346531-2c1855c3-362d-40f9-93d9-6ab38279a275.png)

If your main project file is not of package "main" and is called via a seperate "main" package file, then you don't need to make changes to your project.

Once you have made the necessary changes to your file, if needed, change the function called at line 30 of the goreloaded test file to the main project function if it is different to "GoReloaded()".

Finally to run the test, in a terminal cd to the project folder and run the following command (replacing FILE with your main project file): <br/>
go test FILE goreloaded_test.go

![test_command](https://user-images.githubusercontent.com/81628708/137346962-c5078173-f5fe-4a4d-aed2-2fbb70518d55.png)

If your project passed the test the output should be as follows(with a varying amount of time taken):

![test_passed](https://user-images.githubusercontent.com/81628708/137347384-c2462211-4a52-4d6f-8b3a-3bf28db1eae2.png)

If it did not pass it should be like so:

![test_failed](https://user-images.githubusercontent.com/81628708/137347930-72b43d11-5925-4dab-ab57-94c4bc02ab33.png)
