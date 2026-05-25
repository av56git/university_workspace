import PyPDF2
import sys

def test_pypdf(filepath):
    try:
        reader = PyPDF2.PdfReader(filepath)
        page = reader.pages[0]
        text = page.extract_text()
        print("--- EXTRACTED TEXT ---")
        print(text)
    except Exception as e:
        print("Error:", e)

if __name__ == "__main__":
    test_pypdf(r"C:\Users\daca\.gemini\antigravity\scratch\cyber_ai_workspace\docs\academic_courses\Anno_1\Programmazione\esercizi\Lab 3 - Selezione.pdf")
