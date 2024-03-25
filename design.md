# Design

## API
- GET /
    - root page / sign up
- GET /words/:uid/list
    - list words known by baby uid
- GET /words/:uid/check/:word
    - check whether word is known to baby / return word info
- POST /words/:uid/add/:word
    - add word to known words and return info (# words total)
- DELETE /words/:uid/delete/:word
    - delete word and renumber all later words (nice to have)
- POST /baby
    - Create a new baby (sign up)


## Tables


**babies**

- id
- slug
- name
- birth_date
- created_at


**words**

- id
- baby_id
- word
- number
- learned_date
- created_at

