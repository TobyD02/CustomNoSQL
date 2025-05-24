# TODO

## Collections

- [ ] collections shouldn't be stored in their own file - there should be a collections file that is loaded, where each
  collection stores an array of record IDs and index ids
- [ ] there shouldn't be a load collection endpoint - instead it should be a getAllCollections endpoint

## Indexes

- [ ] indexes should be stored in an index file - each can have its own name but should be referenced by an ID
- [ ] each index should store an array of tokens - where each token has corresponding record IDs from within its
  collection that contains these tokens

## Records

- [ ] Records should be associated with a collection - when created, generate an ID, process it (tokenized fields and
  insert them into indexes), and then add its ID to its corresponding collection 