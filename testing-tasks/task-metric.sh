#!/bin/bash 

#::ARTR::title=Sample Metric Task
#::ARTR::result-type=metric 

free_mem=$(free | grep -i 'mem' | awk '{print($7)}')

echo "Free Memory: ${free_mem}"
