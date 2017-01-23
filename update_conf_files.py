#!/usr/bin/env python

import shutil
import os

if __name__ == "__main__":
	dir_path = os.path.dirname(os.path.realpath(__file__))
	target_path = os.path.join(dir_path, 'configs', 'git')	
	home = os.path.expanduser('~')	
	git_conf = os.path.join(home, '.gitconfig')
	shutil.copy(git_conf, target_path)
	
