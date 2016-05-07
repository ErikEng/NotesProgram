# NotesProgram
This program is meant to help you make better and more informative notes, mainly designed towards KTH students.

Writing notes for lectures can be pretty unpleasant. Often they can be ugly and hard to read, other times it's impossible to understand them without the proper context from the lecture. This program is meant to improve these notes with the use of codewords in the text as well as automated features.
The idea is that you write your notes in the text editor of your choice and when you're done with the text you enter it into the NotesProgram which creates a html document of it with both automatic features as well as features triggered by code words in the text.
So you could for example write :slide1: in the text which would later insert the first slide from the lecture at that day, or you could take a photo of the blackboard and write :photo1: in the text which would later put that photo at that spot in the note. The optimal solution would be if the program could download the lecture slides on it’s own if your provide the webpage that they get uploaded to, but it could also be downloaded to a map the program has access to by hand. The photos could in an ideal scenario be synched with for example a dropbox, but the user could simply move the relevant photos by hand as well. 

##Milestones
1 Turn the given file into a html document of a certain template

2 Recognize codewords and respond to them

3  Able to  insert specific pictures/slides/photos  from a given map into the html as a response to codeword

4  Able to download slides from a given webpage on its own

## Possible additional features
1 Integration with other apps such as Anki or Gdocs so that you by codewords can send information to other apps, such as saving the note to your google drive or put text snippets onto Anki cards

2 By a code word such as :PutIn "Proofs.html": so the coming snippet gets copied into another html document in the map. If the file in the PutIn command can't be found, create a file with that name  

