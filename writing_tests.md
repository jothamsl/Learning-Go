# Writing Tests in Golang (or any other language)

* Write a test
* Make the compiler pass the test
* Run the test, see that it fails and check the error message is meainingful
* Write enough code to make the test pass 
* Refactor


On the face of it this may seem tedious but sticking to the feedback loop is important. Not only does it ensure that you have relevant tests, it helps ensure you design good software by refactoring with the safety of tests.

Seeing the test fail is an important check because it also lets you see what the error message looks like. As a developer it can be very hard to work with a codebase when failing tests do not give a clear idea as to what the problem is.

By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.
By not writing tests you are committing to manually checking your code by running your software which breaks your state of flow and you won't be saving yourself any time, especially in the long run.
