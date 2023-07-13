#!/bin/bash

# Укажите полный путь к файлу database.json
database_file="/mnt/c/Users/develor/Desktop/KoodProject/piscine-go/database.json"

# Получить данные о супергерое с id=170
superhero=$(jq '.[] | select(.id == 170)' "$database_file")

# Извлечь необходимую информацию
name=$(jq -r '.name' <<< "$superhero")
power=$(jq -r '.powerstats.power' <<< "$superhero")
gender=$(jq -r '.appearance.gender' <<< "$superhero")

# Записать результаты в файл
> output.txt  # Очистить файл output.txt
printf "%s\n" "$name" >> output.txt
printf "%s\n" "$power" >> output.txt
printf "%s\n" "$gender" >> output.txt

