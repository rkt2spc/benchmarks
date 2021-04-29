from timeit import default_timer as timer
import re
import gen_data

# Build RegExp Pattern
# @param {object} node is a node in the Trie (usually a hashmap {})
def build_regexp_pattern(node):
    # Leaf node
    if '' in node and len(node.keys()) == 1:
        return None

    # Example 1
    # .
    # f a
    #   b
    #   c z
    #   d z
    # => sub_patterns = ['cz', 'dz'] => RegExp = (?:cz|dz)
    # => sub_characters = ['a', 'b'] => RegExp = [ab]
    # => current_valid = False
    #
    # Example 2
    # .
    # f
    #   a
    #   b
    # => current_valid = True (f is a word in the dictionary)
    sub_patterns = []
    sub_characters = []
    current_valid = False
    for k in sorted(node.keys()):
        child = node[k]
        if isinstance(child, dict):
            child_pattern = build_regexp_pattern(child)
            if child_pattern == None: # child is leaf node
                sub_characters.append(re.escape(k))
            else:
                sub_patterns.append(re.escape(k) + child_pattern)
        else:
            current_valid = True

    # Flag if doesn't have any sub patterns
    no_sub_patterns = len(sub_patterns) == 0

    # Example 1:
    # sub_patterns = ['foo', 'bar']
    # sub_characters = ['a', 'b']
    # => sub_patterns = ['foo', 'bar', '[ab]'] => RegExp (?:foo|bar|[ab])
    #
    # Example 2:
    # sub_patterns = ['foo', 'bar']
    # sub_characters = ['a']
    # => sub_patterns = ['foo', 'bar', 'a'] => RegExp = (?:foo|bar|a)
    if len(sub_characters) > 0:
        if len(sub_characters) == 1:
            sub_patterns.append(sub_characters[0])
        else:
            sub_patterns.append('[' + ''.join(sub_characters) + ']')

    # Example 1
    # sub_patterns = ['foo']
    # => RegExp(/foo/)
    #
    # Example 2
    # sub_patterns = ['foo', 'bar']
    # => RegExp = (?:foo|bar)
    if len(sub_patterns) == 1:
        result = sub_patterns[0]
    else:
        result = "(?:" + "|".join(sub_patterns) + ")"

    # Example 1
    # f
    #   a
    # => RegExp = /fa?/
    #
    # Example 2
    # f
    #   a b
    #   c d
    # => RegExp = /(?:(?:ab|cd))?/
    if current_valid:
        if no_sub_patterns:
            result += "?"
        else:
            result = "(?:%s)?" % result

    return result

# Build Trie
# @param {array.<string>} words
def build_trie(words):
    trie = {}
    for word in words:
        node = trie
        for char in word:
            if char not in node:
                node[char] = {}
            node = node[char]
        node[''] = 1
    return trie

def benchmark_optimized_regexp():
    badwords = gen_data.bad_words()
    messages = gen_data.long_paragraphs(1000)
    
    trie = build_trie(badwords)
    pattern = build_regexp_pattern(trie)
    regexp = re.compile(pattern, re.IGNORECASE)

    start = timer()
    for i in range(len(messages)):
        safe_message = regexp.sub(lambda m: '*' * len(m.group(0)), messages[i])
        # print('[', i + 1, '/', len(messages), ']', messages[i], ' => ', safe_message)

    end = timer()
    print('Finished with', len(messages), 'messages of dictionary size', len(badwords), 'in', end - start, 'seconds')

benchmark_optimized_regexp()
