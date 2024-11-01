import os
import re
import fitz  # PyMuPDF

# Регулярное выражение для поиска email
EMAIL_REGEX = r"(?i)\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}\b"

def get_files_from_dir(directory):
    """Возвращает список всех PDF файлов в указанной директории."""
    try:
        return [f for f in os.listdir(directory) if f.endswith('.pdf')]
    except Exception as e:
        print(f"Ошибка при получении списка файлов: {e}")
        return []

def read_pdf_text(path):
    """Читает текст из PDF файла."""
    try:
        with fitz.open(path) as pdf_document:
            text = ""
            for page_num in range(pdf_document.page_count):
                page = pdf_document[page_num]
                text += page.get_text()
            return text
    except Exception as e:
        print(f"Ошибка при чтении PDF '{path}': {e}")
        return ""

def extract_emails(text):
    """Извлекает все email адреса из текста."""
    return re.findall(EMAIL_REGEX, text)

def main():
    pdf_directory = "../pdfs"
    output_file_path = "output.txt"
    files = get_files_from_dir(pdf_directory)

    print(f"Всего найдено файлов: {len(files)}")
    successful_readings_counter = 0
    unique_emails = set()  # Множество для хранения уникальных email адресов

    for file_name in files:
        pdf_path = os.path.join(pdf_directory, file_name)
        text = read_pdf_text(pdf_path)
        if not text:
            print(f"Пропуск файла из-за ошибки: {file_name}")
            continue

        successful_readings_counter += 1
        emails = extract_emails(text)

        # Добавляем найденные email в множество уникальных
        unique_emails.update(emails)

    # Записываем уникальные email адреса в выходной файл
    with open(output_file_path, 'w', encoding='utf-8') as output_file:
        for email in unique_emails:
            output_file.write(email + "\n")

    print(f"Всего успешно прочитанных файлов: {successful_readings_counter}")
    print(f"Общее количество уникальных email: {len(unique_emails)}")

# Запуск основной функции
if __name__ == "__main__":
    main()
