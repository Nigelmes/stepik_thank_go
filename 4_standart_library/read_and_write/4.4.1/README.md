🤔 Строки из файла

Реализуйте функцию readLines(), которая читает все строки из файла и возвращает их в виде среза:

func readLines(name string) ([]string, error)
В файле строки разделены символом перевода строки \n, а в срезе должны быть уже без него. Если в файле последняя строка пустая (как принято в линуксе), она не попадает в срез.

Если файла не существует или его не удалось прочитать, возвращается ошибка (тип и текст на ваше усмотрение).