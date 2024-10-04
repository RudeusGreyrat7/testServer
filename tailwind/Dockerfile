FROM node:latest

WORKDIR /tailwind

RUN npm init -y && \
    npm install tailwindcss && \
    npx tailwindcss init

CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll

# Пояснение строк кода:

# 	1.	FROM node:latest:
# 	•	Это базовый образ, который Docker использует для сборки вашего контейнера. В данном случае это образ Node.js с последней доступной версией. Он необходим, чтобы иметь все необходимые инструменты для работы с Node.js и npm.
# 	2.	WORKDIR /tailwind:
# 	•	Эта строка задаёт рабочую директорию /tailwind внутри контейнера. Все последующие команды будут выполняться внутри этой директории. Если папка не существует, она будет создана.
# 	3.	RUN npm init -t && npm install tailwindcss && npx tailwindcss init:
# 	•	npm init -t: команда инициализирует новый проект npm с настройками по умолчанию (параметр -t означает, что не будут задаваться вопросы при инициализации).
# 	•	npm install tailwindcss: устанавливает Tailwind CSS как зависимость в проекте.
# 	•	npx tailwindcss init: эта команда с помощью npx инициализирует файл конфигурации Tailwind CSS (tailwind.config.js), который нужен для настройки Tailwind в проекте.
# 	4.	CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll:
# 	•	npx tailwindcss: запускает Tailwind CSS через npx, который временно использует локальные зависимости, не требуя глобальной установки.
# 	•	-c /src/tailwind.config.js: указывает на кастомный файл конфигурации Tailwind CSS, который находится по пути /src/tailwind.config.js.
# 	•	-i /src/styles.css: задаёт входной CSS-файл (styles.css), в котором содержатся директивы для Tailwind (@tailwind base;, @tailwind components;, @tailwind utilities;).
# 	•	-o /dst/styles.css: задаёт путь для выходного файла, куда Tailwind сгенерирует финальный CSS. В данном случае он будет сохранён в директории /dst как styles.css.
# 	•	--watch: команда запускает Tailwind в режиме “watch”, при котором он будет автоматически отслеживать изменения в файлах и пересобирать CSS при каждом изменении.
# 	•	--poll: этот параметр используется для контроля изменений файлов через периодические опросы системы (polling). Полезно, если файловая система или среда (например, Docker) не поддерживает события изменения файлов в реальном времени.

# Общая логика:

# 	•	Вы создаёте образ Docker с установленным Tailwind CSS и настройками для его автоматической сборки. В контейнере будет инициализирован проект Node.js, установлены зависимости и настроен процесс генерации CSS из Tailwind с автообновлением при изменении исходных файлов.