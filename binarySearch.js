const binarySearch = (list, item) =>{

    let low = 0;
    let high = list.length - 1;

    while (low <= high){
        let middle = Math.floor((low + high) /2);
        let guessTry = list[middle];
        if(guessTry == item){
            return middle;
        }
        if(guessTry > item){
            high = middle - 1;
        }else{
            low = middle + 1;
        }
    }

    return null;
}

const exampleList = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

console.log(binarySearch(exampleList, 10));