from timeit import default_timer as timer
import re
import gen_data

def benchmark_regexp():
    badwords = gen_data.bad_words()
    messages = gen_data.long_paragraphs(1000)

    pattern = '|'.join(map(re.escape, badwords))
    regexp = re.compile(pattern, re.IGNORECASE)

    start = timer()
    for i in range(len(messages)):
        safe_message = regexp.sub(lambda m: '*' * len(m.group(0)), messages[i])
        # print('[', i + 1, '/', len(messages), ']', messages[i], ' => ', safe_message)

    end = timer()
    print('Finished with', len(messages), 'messages of dictionary size', len(badwords), 'in', end - start, 'seconds')
        
benchmark_regexp()