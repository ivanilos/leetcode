class Solution:
    def simplifyPath(self, path: str) -> str:
        splitted_path = path.split('/')
        
        new_path = []
        for directory in splitted_path:
            if directory == '..':
                if len(new_path) > 0:
                    new_path.pop()
            elif directory not in ('.', ''):
                print("here dir = {}".format(directory))
                new_path.append(directory)
        
        return '/' + '/'.join(new_path)
        