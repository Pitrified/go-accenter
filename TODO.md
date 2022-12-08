# TODO and ideas

1. IPA pronunciation
1. Filter out words very similar to English.
1. Optional accent filter.
1. Cluster by categories.
1. Show all senses.
1. Show number of letters.
1. Show which letter is accented.
1. Remove surnames and first names.
   Filter `categories` similar to `French surnames`, `French given names`, `French female given names`.
1. Button to blacklist a word.
1. Button to show more info on the word (categories, form_of, ...)
1. Buttons to skip the word or request some hints.
1. Loading the records asynchronously and showing a progress bar would be a neat exercise.
1. View with what is basically a dictionary that shows all the words with glosses/freq/errors? Sortable by error?
1. Option to give a lot more weight to words with errors.
1. Some way to backup the pretend database.
1. While the user is typing, find the new word.
   Use the old weights who cares.
   So we can do it in the most inefficient way and the user will not notice.
1. When inserting a correct word that had errors,
   divide the number of errors by some factor,
   rather than removing one.
   If we boost the wrong words a lot you risk to make a typo
   which will keep increasing the word weight.
   When we learn the word we want to quickly reduce it.
