# Author: Xujianxu
# -*- coding: utf-8 -*-
# coding=utf-8
import xlrd
import xlwt
import os
import time
import smtplib
from datetime import date,datetime
from xlrd import xldate_as_tuple

def read_excel():
    # 打开文件
    
    path = "C:\collectinfo"
    files = os.listdir(path)
    row_list = []


    for file in files:
        #print (file)
        #print (path)
        filepath = path+ '\\' + file
        #print (filepath)
        workbook = xlrd.open_workbook(filepath)

        sheet2 = workbook.sheet_by_index(0) 
        nrows = sheet2.nrows
#        print (sheet2.name,sheet2.nrows,sheet2.ncols)
     
        
        ##把数据放入一个二维数组
        for i in range(0,sheet2.nrows):
            row = sheet2.row_values(i)
            row_list.append(row)
        #print(row_list)
    return row_list
            
 #           for m in range(0,sheet2.ncols):
 #               sheet1.write(i,m,row[m])
                
            


    
def dedup():
    work_path = r"C:\scripts\total.xls"
    work_book = xlrd.open_workbook(work_path)

    sheet = work_book.sheet_by_index(0) 
    n_rows = sheet.nrows
    
    print (sheet.name,sheet.nrows,sheet.ncols)
    #write to new file 
    new_wbk = xlwt.Workbook()
    new_sheet1 = new_wbk.add_sheet('sheet1')  
    
    for n in range(0,sheet.ncols):
        #print(sheet.cell_value(0,n))
        new_sheet1.write(0,n,sheet.cell_value(0,n))
    
    t=1
    for i in range(1,sheet.nrows):
        #row = sheet.row_values(i)
        #print(row)
        #print(sheet.cell(i,0)) 
        cellvalue = sheet.cell(i,0).value
        #print(cellvalue)
        if cellvalue != "序号" and cellvalue != "" :
            print(t)
            print(cellvalue)
            #print("need to delete") 
            for s in range(0,sheet.ncols):
                new_sheet1.write(t,s,sheet.cell_value(i,s))
            t=t+1
    finalfile = "final.xls"
    new_wbk.save(finalfile)
                
            #sheet.deleteRow(sheet,i)
            #work_book.save()
	
    # 获取单元格内容
    #print sheet2.cell(1,0).value.encode('utf-8')
    #print sheet2.cell_value(1,0).encode('utf-8')
    #print sheet2.row(1)[0].value.encode('utf-8')
    
    # 获取单元格内容的数据类型
    #print sheet2.cell(1,0).ctype

if __name__ == '__main__':
    rows = read_excel()
    print(len(rows))
    #print(len(rows[0]))
    #print(row)

    wbk = xlwt.Workbook()
    sheet1 = wbk.add_sheet('sheet1')
    for m in range(0,len(rows)):
        for i in range(0,len(rows[m])):
            sheet1.write(m,i,rows[m][i])
    newfile = "total.xls"
    wbk.save(newfile)
    dedup()
    os.remove(newfile)

    
