# create for python learning
# get the size for whole folder in linux
# author : Xujianxu
import os
import sys
import argparse

def get_parser():
    parser = argparse.ArgumentParser(description="Please input the full folder path")
    parser.add_argument('--path', help="choose a folder to get the size")
    args = parser.parse_args()
    print parser
    return parser

def get_size(folder):
    dir_size = 0
    print folder
    for (path, dirs, files) in os.walk(folder):
        for file in files:
            filename = os.path.join(path, file)
#            print path, file
            dir_size += os.path.getsize(filename)
    print dir_size
    kb_size = float(dir_size)/1024
    mb_size = float(dir_size)/(1024 * 1024)
    gb_size = float(dir_size)/(1024 * 1024 *1024)
    print str(kb_size)+str('kb')
    print str(mb_size)+str("MB")
    print str(gb_size)+str("GB")


if __name__ == '__main__':
    parser = get_parser()
    args = vars(parser.parse_args())
    folder = args['path']
    get_size(folder)
