#! /usr/bin/env python3
import sys
import os
from janome.tokenizer import Tokenizer

tokenizer = Tokenizer()

def _is_main_word(token):
    if '名詞' in token.part_of_speech:
        return True
    if '動詞' in token.part_of_speech and '自立' in token.part_of_speech:
        return True
    return False

def _concat(tokens):
    words = [t.surface for t in tokens]
    return ''.join(words)

def wrap(line):
    tokens = tokenizer.tokenize(line)
    results = []
    i = 0
    while i < len(tokens):
        start = i
        i += 1
        if '名詞' in tokens[start].part_of_speech:
            while i < len(tokens) and '名詞' in tokens[i].part_of_speech:
                i += 1
        while i < len(tokens) and not _is_main_word(tokens[i]):
            i += 1
        results.append(_concat(tokens[start:i]))
    words = ['<span class="wordwrap">{}</span>'.format(result) for result in results]
    return ''.join(words)

if __name__ == '__main__':
    for line in sys.stdin.readlines():
        if line.strip() == "":
            print()
            continue
        print('<p>{}</p>'.format(wrap(line)))
