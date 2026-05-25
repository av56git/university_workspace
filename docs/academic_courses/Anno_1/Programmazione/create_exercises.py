import os
import re
import glob
import PyPDF2

base_dir = r"C:\Users\daca\.gemini\antigravity\scratch\cyber_ai_workspace\docs\academic_courses\Anno_1\Programmazione"
pdf_dir = os.path.join(base_dir, "esercizi")
out_dir = os.path.join(base_dir, "eseguendo")

def sanitize_filename(name):
    # Replace unicode replacement chars
    name = name.replace("", "")
    # Remove invalid Windows chars
    name = re.sub(r'[<>:"/\\|?*]', '', name)
    return name.strip()

def process_pdfs():
    pdfs = glob.glob(os.path.join(pdf_dir, "*.pdf"))
    for pdf_path in pdfs:
        pdf_name = os.path.basename(pdf_path)
        pdf_basename = os.path.splitext(pdf_name)[0]
        
        try:
            reader = PyPDF2.PdfReader(pdf_path)
            # Try first 2 pages
            text = ""
            for i in range(min(2, len(reader.pages))):
                page_text = reader.pages[i].extract_text()
                if page_text:
                    text += page_text + "\n"
            
            exercises = []
            for line in text.split('\n'):
                line = line.strip()
                # Matches: "1 Qual è l'output? ........... 2"
                match = re.match(r"^\d+\s+(.*?)\.{4,}", line)
                if match:
                    ex_name = match.group(1).strip()
                    exercises.append(ex_name)
            
            if exercises:
                # Create folder
                folder_path = os.path.join(out_dir, pdf_basename)
                os.makedirs(folder_path, exist_ok=True)
                
                # Create files
                for ex in exercises:
                    safe_name = sanitize_filename(ex)
                    file_path = os.path.join(folder_path, safe_name + ".go")
                    # don't overwrite if exists
                    if not os.path.exists(file_path):
                        with open(file_path, 'w', encoding='utf-8') as f:
                            pass
                print(f"Processed {pdf_name}: created {len(exercises)} files.")
            else:
                print(f"No exercises found in TOC for {pdf_name}")
                
        except Exception as e:
            print(f"Error processing {pdf_name}: {e}")

if __name__ == "__main__":
    process_pdfs()
