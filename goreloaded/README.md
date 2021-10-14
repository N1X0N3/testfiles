This test file will check your go-reloaded project against the test cases on the intra audit web page.

The test file needs to be copied to your project folder, the same folder that your main project file is located.

If your main project file is of package "main" and your project is processed in the "main()" function, then you will need to change the name of that function to GoReloaded(), or another name of your choosing (has to start with a capital letter, to make it public), and then create another function called main() where you call the original functions new name.

If your main project file is not of package "main" and is called via a seperate "main" package file, then you don't need to make changes to your project.

Once you have made the necessary changes to your file, if needed, change the function called at line 30 of the goreloaded test file to the main project function if it is different to "GoReloaded()"

Finally to run the test, in a terminal cd to the project folder and run the following command:
go test <"main project file name"> goreloaded_test.go