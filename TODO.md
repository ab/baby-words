# TODO

- [x] Upgrade go to 1.21+ so I can use bingo to install dbmate
  - https://go.dev/dl/
  - https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
  - bingo get -l github.com/amacneil/dbmate/v2

- [x] Basic API operations
    - [x] create baby
    - [x] get baby
    - [x] create word
    - [x] list words

- [ ] Basic UI
    - [x] HTMl for list and create words
    - [ ] New index page, HTML for create baby
    - [ ] JS to avoid duplicate words
    - [ ] Party for first word
    - [ ] XHR for creating words

- [ ] Deploy to fly.io
    - [ ] Decide whether to use LiteFS vs. just a single node with sqlite
    - [ ] Launch app

- [ ] Import existing data
    - [ ] Add existing words from my spreadsheet, backdating estimated `learned_date` smoothly over each month
        - `curl -i -d "word=omg&learned_date=2024-01-01" -X POST http://localhost:8080/words/<slug>/add`

- [ ] `client_info` logging

- [ ] View-only link?

- [ ] Basic stats
    - [ ] Your baby knows N words, is M days old

- [ ] Chart words learned over time
