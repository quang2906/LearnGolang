package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Job       string `json:"job"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Salary    int    `json:"salary"`
	BirthDate string `json:"birthdate"`
}

func GroupPeopleByCity(persons []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range persons {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

func CountPeopleByCity(persons []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range persons {
		result[person.City]++
	}
	return result
}

func GroupPeopleByJob(persons []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range persons {
		result[person.Job]++
	}
	return result
}

func GetPeopleByJob(persons []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range persons {
		result[person.Job] = append(result[person.Job], person)
	}
	return result
}

func Top5JobsByNumber(persons []Person) (result []string) {
	jobs := GroupPeopleByJob(persons)
	type kv struct {
		Key   string
		Value int
	}

	var jobs_arr []kv
	for k, v := range jobs {
		jobs_arr = append(jobs_arr, kv{k, v})
	}

	sort.Slice(jobs_arr, func(i, j int) bool {
		return jobs_arr[i].Value > jobs_arr[j].Value
	})

	for i := 0; i < 5; i++ {
		result = append(result, jobs_arr[i].Key)
	}
	return result
}

func Top5CitiesByNumber(persons []Person) (result []string) {
	cities := CountPeopleByCity(persons)
	type kv struct {
		Key   string
		Value int
	}

	var cities_arr []kv
	for k, v := range cities {
		cities_arr = append(cities_arr, kv{k, v})
	}

	sort.Slice(cities_arr, func(i, j int) bool {
		return cities_arr[i].Value > cities_arr[j].Value
	})

	for i := 0; i < 5; i++ {
		result = append(result, cities_arr[i].Key)
	}
	return result
}

func TopJobByNumberInEachCity(persons []Person) (result map[string]string) {
	result = make(map[string]string)
	cities := GroupPeopleByCity(persons)
	for city, personsOfCity := range cities {
		top5Jobs := Top5JobsByNumber(personsOfCity)
		result[city] = top5Jobs[0]
	}
	return result
}

func AverageSalaryByJob(persons []Person) (result map[string]int) {
	result = make(map[string]int)
	jobs := GetPeopleByJob(persons)
	for job, personsOfJob := range jobs {
		sumSalary := 0
		for _, person := range personsOfJob {
			sumSalary += person.Salary
		}
		result[job] = sumSalary / len(personsOfJob)
	}
	return result
}

func AverageSalaryByCity(persons []Person) (result map[string]int) {
	result = make(map[string]int)
	cities := GroupPeopleByCity(persons)
	for city, personsOfCity := range cities {
		sumSalary := 0
		for _, person := range personsOfCity {
			sumSalary += person.Salary
		}
		result[city] = sumSalary / len(personsOfCity)
	}
	return result
}

func FiveCitiesHasTopAverageSalary(persons []Person) (result []string) {
	cities := AverageSalaryByCity(persons)
	type kv struct {
		Key   string
		Value int
	}

	var cities_arr []kv
	for k, v := range cities {
		cities_arr = append(cities_arr, kv{k, v})
	}

	sort.Slice(cities_arr, func(i, j int) bool {
		return cities_arr[i].Value > cities_arr[j].Value
	})

	for i := 0; i < 5; i++ {
		result = append(result, cities_arr[i].Key)
	}
	return result
}

func FiveCitiesHasTopSalaryForDeveloper(persons []Person) (result []string) {
	avgSalaryOfDeveloperInEachCity := make(map[string]int)
	cities := GroupPeopleByCity(persons)
	for city, personsOfCity := range cities {
		avgSalaryOfJob := AverageSalaryByJob(personsOfCity)
		avgSalaryOfDeveloperInEachCity[city] = avgSalaryOfJob["developer"]
	}
	type kv struct {
		Key   string
		Value int
	}

	var cities_arr []kv
	for k, v := range avgSalaryOfDeveloperInEachCity {
		cities_arr = append(cities_arr, kv{k, v})
	}

	sort.Slice(cities_arr, func(i, j int) bool {
		return cities_arr[i].Value > cities_arr[j].Value
	})

	for i := 0; i < 5; i++ {
		result = append(result, cities_arr[i].Key)
	}
	return result
}

func GetAge(age string) (result int) {
	arr := strings.Split(age, "-")
	year, _ := strconv.Atoi(arr[0])
	//month, _ := strconv.Atoi(arr[1])
	//day, _ := strconv.Atoi(arr[2])
	result = 2021 - year
	return result
}

func AverageAgePerJob(persons []Person) (result map[string]float64) {
	result = make(map[string]float64)
	jobs := GetPeopleByJob(persons)
	for job, personsOfJob := range jobs {
		sumAge := 0.0
		for _, person := range personsOfJob {
			sumAge += (float64)(GetAge(person.BirthDate))
		}
		result[job] = sumAge / (float64)(len(personsOfJob))
	}
	return result
}

func AverageAgePerCity(persons []Person) (result map[string]float64) {
	result = make(map[string]float64)
	cities := GroupPeopleByCity(persons)
	for city, personsOfCity := range cities {
		sumAge := 0.0
		for _, person := range personsOfCity {
			sumAge += (float64)(GetAge(person.BirthDate))
		}
		result[city] = sumAge / (float64)(len(personsOfCity))
	}
	return result
}

func main() {
	jsonFile, err := os.Open("person.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var persons []Person
	json.Unmarshal(byteValue, &persons)

	fmt.Println("-----Gom tất cả những người trong cùng một thành phố-----")
	fmt.Println(GroupPeopleByCity(persons))

	fmt.Println("-----Các nghề nghiệp và số người làm-----")
	fmt.Println(GroupPeopleByJob(persons))

	fmt.Println("-----5 nghề có nhiều người làm nhất-----")
	fmt.Println(Top5JobsByNumber(persons))

	fmt.Println("-----5 thành phố có nhiều người ở nhất-----")
	fmt.Println(Top5CitiesByNumber(persons))

	fmt.Println("-----Nghề được làm nhiều nhất trong mỗi thành phố-----")
	fmt.Println(TopJobByNumberInEachCity(persons))

	fmt.Println("-----Mức lương trung bình của mỗi nghề-----")
	fmt.Println(AverageSalaryByJob(persons))

	fmt.Println("-----5 thành phố có mức lương trung bình cao nhất-----")
	fmt.Println(FiveCitiesHasTopAverageSalary(persons))

	fmt.Println("-----5 thành phố có mức lương trung bình của developer cao nhất----- ")
	fmt.Println(FiveCitiesHasTopSalaryForDeveloper(persons))

	fmt.Println("-----Tuổi trung bình của từng nghề nghiệp-----")
	fmt.Println(AverageAgePerJob(persons))

	fmt.Println("-----Tuổi trung bình ở từng thành phố-----")
	fmt.Println(AverageAgePerCity(persons))
}
