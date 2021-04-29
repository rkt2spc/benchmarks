import json

def long_paragraphs(n = 1):
    paragraph = """
Intrigued by its descriptions of meditation and yoga and how our spirits can transcend our material body, he says the encounter made an impression him as a teenager. Looking back, he recalls that he felt he had found a kind of truth. “This is what I have been looking for,” he thought.
    
Edwards, who is now known as Jai Nitai Dasa, became a vegetarian, took up yoga, and started meditating. Back then, that was as far as he was willing to go.
    
Fifteen years later, he was working in investment banking in London. This was 1990s Britain, post–Margaret Thatcher, and the UK’s financial district was running with minimum government oversight and at the mercy of the free market. A lot of people were making a lot of money, fast.
"""

    return [paragraph] * n

def real_messages(n = None):
    messages = list(filter(None, json.load(open('messages.json'))))
    return messages if n == None else messages[:min(len(messages), min(n, 0))]

def bad_words(n = None):
    badwords = json.load(open('badwords.json'))
    return badwords if n == None else badwords[:min(len(badwords), min(n, 0))]
