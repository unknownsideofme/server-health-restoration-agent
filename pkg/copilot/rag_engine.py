#!/usr/bin/env python3
"""
AirGap Local Offline RAG Engine
Maintains an air-gapped vector/document index over internal runbooks, topology metadata, and incident history.
No external network or cloud dependencies.
"""

import glob
import os
import re
from typing import List, Dict


class LocalRAGEngine:
    def __init__(self, docs_dir: str = "docs/runbooks"):
        self.docs_dir = docs_dir
        self.documents: List[Dict[str, str]] = []
        self.load_documents()

    def load_documents(self):
        self.documents.clear()
        if not os.path.exists(self.docs_dir):
            return

        for filepath in glob.glob(os.path.join(self.docs_dir, "*.md")):
            filename = os.path.basename(filepath)
            try:
                with open(filepath, "r", encoding="utf-8") as f:
                    content = f.read()
                    self.documents.append({
                        "id": filename,
                        "title": filename.replace(".md", "").replace("_", " ").title(),
                        "path": filepath,
                        "content": content,
                    })
            except Exception:
                pass

    def retrieve_context(self, query: str, top_k: int = 2) -> List[dict]:
        query_words = set(re.findall(r"\w+", query.lower()))
        scored_docs = []

        for doc in self.documents:
            doc_words = set(re.findall(r"\w+", doc["content"].lower()))
            overlap = len(query_words.intersection(doc_words))
            scored_docs.append((overlap, doc))

        scored_docs.sort(key=lambda x: x[0], reverse=True)
        results = [doc for score, doc in scored_docs[:top_k] if score > 0]

        if not results and self.documents:
            results = [self.documents[0]]

        return results
