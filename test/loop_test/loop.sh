#!/bin/bash


#cityFUN(){
#     echo "city: $1"
#     ./city_info.sh $1
#
#}

business_areaFUN(){
     ./business_area.sh $1 $2

}

business_districtFUN(){
     ./business_district.sh $1 $2
}
buildingFUN(){
     ./building.sh $1 $2
}
fridgeFUN(){
     echo "fridge:"
     echo "	area: $4"
     echo "		district: $3"
     echo "			building: $2"
     echo "				   id: $1"
     echo "					fridge_storage x 9"
     ./fridge.sh $1 $2 $3 $4 
}
ordersFUN(){
     #echo "orders: $1"
     ./orders.sh $1 $2
}
fridge_storageFUN(){
     ./fridge_storage.sh $1 $2
     #echo "fridge_storage: $1"
}
num=10
#for ((city=1; city<num; ++city))
#do
#    cityFUN $city
    for ((area=1; area<num; ++area))
    do
	let a=$area*1000
        business_areaFUN $a $city 
        for ((district=1; district<num; ++district))
        do
            let d=$district*100+$area*1000
	    let a=$area*1000
            business_districtFUN $d $a
            for ((building=1; building<num; ++building))
            do
		let b=${building}*10+${district}*100+${area}*1000
		let d=${district}*100+${area}*1000
                buildingFUN $b $d
                for ((fridge=1; fridge<num; ++fridge))
                do
		    let f=${fridge}+${building}*10+${district}*100+${area}*1000
		    let b=${building}*10+${district}*100+${area}*1000
		    let d=district*100+area*1000
		    let a=${area}*1000

                    fridgeFUN ${f} ${b} ${d} ${a}
 		    for ((orders=1; orders<num*2; ++orders))
                    do
		        let f=${fridge}+${building}*10+${district}*100+${area}*1000
                        ordersFUN MJ${f} ${f}
	            done
		    for ((storage=1; storage<num; ++storage))
		    do
		        let f=${fridge}+${building}*10+${district}*100+${area}*1000
		        fridge_storageFUN ${f} ${storage}
		    done

                done
            done
        done
     done
#done
