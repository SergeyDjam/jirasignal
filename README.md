# jirasignal
The alarm system about tasks in Jira on GO

1. ```git clone https://github.com/SergeyDjam/jirasignal```
2. ```cd jirasignal```
3. ```mv .env.example .env```
4. ```vi .env ```
    заполняем все параметры
5. ```go run main.go``` или если хочется отдельный бинарничек ```go build && ./jirasignal``` или ```go get github.com/SergeyDjam/jirasignal && go/bin/jirasignal```


Эта программа переделка вот этой прогаммы: https://github.com/Osuka42g/JiraNotifier

GPL