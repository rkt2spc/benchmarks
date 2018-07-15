from timeit import default_timer as timer
import re
import gen_data

def replace(match, badwords_dictionary):
    word = match.group(0)
    if word in badwords_dictionary:
        return '*' * len(word)

    return word

def benchmark_simple():
    badwords = set(gen_data.bad_words()) # hashmap
    messages = gen_data.long_paragraphs(1000)

    regexp = re.compile('\w+', re.IGNORECASE)

    start = timer()
    for i in range(len(messages)):
        safe_message = regexp.sub(lambda m: replace(m, badwords), messages[i])
        # print('[', i + 1, '/', len(messages), ']', messages[i], ' => ', safe_message)

    end = timer()
    print('Finished with', len(messages), 'messages of dictionary size', len(badwords), 'in', end - start, 'seconds')
        
benchmark_simple()