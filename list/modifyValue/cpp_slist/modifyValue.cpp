#include <iostream>
#include <fstream>
#include <vector>
#include "modifyValue.h"


void print_list(const slist<int> &argu){
    if(!argu.is_sNil()){
      std::cout<<argu.as_sCons().p1()<<" ";
      print_list(argu.as_sCons().p2());
    }
    if(argu.is_sNil()){
        std::cout << "nil" << std::endl;
    }
}

int main(int argc, char **argv) {
    if (argc != 3) {
        std::cout << "usage: ./main <size> <target>\n"; 
        exit(1);
    }
    uint64_t len = strtol(argv[1],NULL,10);
    uint64_t tar = strtol(argv[2],NULL,10);
    std::cout << "len: " << len << "   tar: " << tar << std::endl;


    slist<int> newList = slist<int>::sNil();

    // 输出读取的整数
    // std::cout << "从文件中读取的整数:" << std::endl;
    for (size_t i = 1; i <= len; ++i) {
        newList = slist<int>::sCons(i, std::move(newList));
        // std::cout << numbers[i] << " ";
    }
    // print_list(newList);
    // std::cout << std::endl;

    struct timespec start, end;
    slist<int> copy_list;
    if(tar == 0){
        timespec_get(&start, TIME_UTC);
        copy_list = ModifyValue(1, 0, newList);
        timespec_get(&end, TIME_UTC);
    }else{
        timespec_get(&start, TIME_UTC);
        copy_list = ModifyValue(1, 0, newList);
        timespec_get(&end, TIME_UTC);
    }

    long long elapsed = (end.tv_sec - start.tv_sec) * 1000000000LL + (end.tv_nsec - start.tv_nsec);
    // print_list(copy_list);
    
    // 输出结果
    std::cout <<"  Time: " << elapsed <<" ns"<<std::endl;

    return 0;
}