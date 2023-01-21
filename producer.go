package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// // Go-routine to handle message delivery reports and
	// // possibly other event types (errors, stats, etc)
	// go func() {
	// 	for e := range p.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
	// 					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
	// 			}
	// 		}
	// 	}
	// }()
	min := 0
	max := len(nms)

	for range time.Tick(time.Second * 2) {
		producemsg(nms[rand.Intn(max-min)+min])
	}

}

func producemsg(msg string) {
	m := kafka.ConfigMap{}
	m["bootstrap.servers"] = "localhost:9092"

	topic := "anshu"
	p, err := kafka.NewProducer(&m)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("name"),
		Value:          []byte(msg),
	}, nil)
	// Wait for all messages to be delivered
	p.Flush(100)
	p.Close()
}

var nms []string = []string{
	"Aadi",
	"Aarav",
	"Aarnav",
	"Aarush",
	"Aayush",
	"Abdul",
	"Abeer",
	"Abhimanyu",
	"Abhiram",
	"Aditya",
	"Advaith",
	"Advay",
	"Advik",
	"Agastya",
	"Akshay",
	"Amol",
	"Anay",
	"Anirudh",
	"Anmol",
	"Ansh",
	"Arin",
	"Arjun",
	"Aryan",
	"Atharv",
	"Avi",
	"Ayaan",
	"Ayush",
	"Ayushman",
	"Azaan",
	"Azad",
	"Brijesh",
	"Bachittar",
	"Bahadurjit",
	"Bakhshi",
	"Balendra",
	"Balhaar",
	"Baljiwan",
	"Balvan",
	"Balveer",
	"Banjeet",
	"Chaitanya",
	"Chakradev",
	"Chakradhar",
	"Champak",
	"Chanakya",
	"Chandran",
	"Chandresh",
	"Charan",
	"Chatresh",
	"Chatura",
	"Daksh",
	"Darsh",
	"Dev",
	"Devansh",
	"Dhruv",
	"Dakshesh",
	"Dalbir",
	"Darpan",
	"Ekansh",
	"Ekalinga",
	"Ekapad",
	"Ekavir",
	"Ekaraj",
	"Ekbal",
	"Farhan",
	"Falan",
	"Faqid",
	"Faraj",
	"Faras",
	"Fitan",
	"Fariq",
	"Faris",
	"Fiyaz",
	"Frado",
	"Gautam",
	"Gagan",
	"Gaurang",
	"Girik",
	"Girindra",
	"Girish",
	"Gopal",
	"Gaurav",
	"Gunbir",
	"Guneet",
	"Harsh",
	"Harshil",
	"Hredhaan",
	"Hardik",
	"Harish",
	"Hritik",
	"Hitesh",
	"Hemang",
	"Isaac",
	"Ishaan",
	"Imaran",
	"Indrajit",
	"Ikbal",
	"Ishwar",
	"Jason",
	"Jagdish",
	"Jagat",
	"Jatin",
	"Jai",
	"Jairaj",
	"Jeet",
	"Kabir",
	"Kalpit",
	"Karan",
	"Kiaan",
	"Krish",
	"Krishna",
	"Laksh",
	"Lucky",
	"Lakshit",
	"Lohit",
	"Laban",
	"Manan",
	"Mohammed",
	"Madhav",
	"Mitesh",
	"Maanas",
	"Manbir",
	"Maanav",
	"Manthan",
	"Nachiket",
	"Naksh",
	"Nakul",
	"Neel",
	"Nakul",
	"Naveen",
	"Nihal",
	"Nitesh",
	"Om",
	"Ojas",
	"Omkaar",
	"Onkar",
	"Onveer",
	"Orinder",
	"Parth",
	"Pranav",
	"Praneel",
	"Pranit",
	"Pratyush",
	"Qabil",
	"Qadim",
	"Qarin",
	"Qasim",
	"Rachit",
	"Raghav",
	"Ranbir",
	"Ranveer",
	"Rayaan",
	"Rehaan",
	"Reyansh",
	"Rishi",
	"Rohan",
	"Ronith",
	"Rudra",
	"Rushil",
	"Ryan",
	"Sai",
	"Saksham",
	"Samaksh",
	"Samar",
	"Samarth",
	"Samesh",
	"Sarthak",
	"Sathvik",
	"Shaurya",
	"Shivansh",
	"Siddharth",
	"Tejas",
	"Tanay",
	"Tanish",
	"Tarak",
	"Teerth",
	"Tanveer",
	"Udant",
	"Udarsh",
	"Utkarsh",
	"Umang",
	"Upkaar",
	"Vedant",
	"Veer",
	"Vihaan",
	"Viraj",
	"Vivaan",
	"Wahab",
	"Wazir",
	"Warinder",
	"Warjas",
	"Wriddhish",
	"Wridesh",
	"Yash",
	"Yug",
	"Yatin",
	"Yuvraj",
	"Yagnesh",
	"Yatan",
	"Zayan",
	"Zaid",
	"Zayyan",
	"Zashil",
	"Zehaan",
	"Aadrik",
	"Aarv",
	"Aachman",
	"Aadhyatm",
	"Aadhavan",
	"Aadarsh",
	"Aadamya",
	"Baidyanath",
	"Bala Shankar",
	"Balamurugan",
	"Bakhtawar",
	"Bagira",
	"Badrinarayanan",
	"Bhargava",
	"Charish",
	"Chaanakya",
	"Chaidya",
	"Chaital",
	"Chinmayee",
	"Dignesh",
	"Dahak",
	"Daityakarya",
	"Daivya",
	"Ekachith",
	"Evyavan",
	"Fanibhushan",
	"Fravash",
	"Ganadhakshya",
	"Gandhamadhana",
	"Garuda",
	"Geetham",
	"Geethik",
	"Govindaram",
	"Hansaraj",
	"Hareshwar",
	"Harikant",
	"Ilanthirayan",
	"Inbanathan",
	"Kalaparan",
	"Kshatragna",
	"Mullinti",
	"Nayanesh",
	"Nihant",
	"Prajith",
	"Reshvind",
	"Rajasekaran",
	"Thanvye",
	"Srivasthav",
	"Sridhara",
	"Vishu",
	"Aadithyakethu",
	"Aadhirai",
	"Aakalp",
	"Aanandit",
	"Abhinivesh",
	"Bala Subramani",
	"Bhagyanandana",
	"Bhooteshwara",
	"Bindusar",
	"Batnasiddhikara",
	"Chintransh",
	"Chithravarma",
	"Chitrasen",
	"Chithrakundhala",
	"Danavendra",
	"Danush",
	"Dasaradh",
	"Devayan",
	"Girinath",
	"Gangasiruvan",
	"Gouthaman",
	"Indradyumm",
	"Jayadeva",
	"Jagadisha",
	"Jaiprakash",
	"Kaanishk",
	"Kabalikruta",
	"Kalaivanan",
	"Lakshmipriya",
	"Lakshiminath",
	"Lankapuravidahaka",
	"Lathesh",
	"Madeshwaran",
	"Madhusudan",
	"Mahadhyuta",
	"Mahakaleshwar",
	"Nadapratithishta",
	"Navamani",
	"Nakulesh",
	"Oshin",
	"Prasannatmane",
	"Parasmaijyotishe",
	"Raaghava",
	"Radhatanaya",
	"Sudharma",
	"Samarpana",
	"Sriharsha",
	"Vishveshwara",
	"Vijayrathna",
	"Venkatesh",
	"Ayushya",
	"Arulappan",
	"Arindham",
	"Akarshan",
	"Aalhad",
	"Agnikumara",
	"Bal Mukund",
	"Balaark",
	"Bajrang",
	"Banaj",
	"Chandrasai",
	"Chellamuthu",
	"Cheranraj",
	"Chidakash",
	"Chellamani",
	"Dhishan",
	"Dhyutidhara",
	"Danavarsh",
	"Dashagreevakulantaka",
	"Dhridhakarmaavu",
	"Eeswar",
	"Elilvendan",
	"Ethiraj",
	"Gaoushik",
	"Garjan",
	"Kaachim",
	"Kaanchanadhwaja",
	"Murlimanohar",
	"Maandavik",
	"Mal Marugan",
	"Naavinya",
	"Nabhanyu",
	"Nadeesh",
	"Nagabhushana",
	"Nageswara",
	"Omkareshwar",
	"Palanichamy",
	"Panchajana",
	"Raajan",
	"Raghukumara",
	"Saavyas",
	"Sabareeshwara",
	"Sacchidananda",
	"Tusya Udarchis",
	"Tatvagyanaprada",
	"Udayasooriyan",
	"Vighnarajendra",
	"Vagadheeksha",
	"Vaijeenath",
	"Venkatesh",
	"Aarsh",
	"Amish",
	"Batuk",
	"Bhavesh",
	"Bhavik",
	"Bhavin",
	"Chintan",
	"Chirayu",
	"Daiwik",
	"Daksh",
	"Darsh",
	"Deval",
	"Dhaval",
	"Dharmesh",
	"Ernesh",
	"Falgun",
	"Fenil",
	"Girish",
	"Hemendra",
	"Herik",
	"Himesh",
	"Ishanyu",
	"Jignesh",
	"Jigar",
	"Jigisu",
	"Joshil",
	"Jatin",
	"Jayesh",
	"Kundan",
	"Kunjan",
	"Kirat",
	"Kirti",
	"Mukund",
	"Munjal",
	"Mitesh",
	"Mahesh",
	"Mahindra",
	"Mansukh",
	"Naman",
	"Naresh",
	"Navin",
	"Neel",
	"Nikesh",
	"Nikunj",
	"Nilesh",
	"Nimish",
	"Paresh",
	"Ramesh",
	"Sudhir",
	"Sunil",
	"Abhinay",
	"Ajinkya",
	"Ashwin",
	"Amey",
	"Bhushan",
	"Chandrakant",
	"Chaitanya",
	"Chinmay",
	"Gajanan",
	"Gandharv",
	"Harshal",
	"Harshavardhan",
	"Lavnik",
	"Mandar",
	"Mirajkar",
	"Mayur",
	"Neeraj",
	"Neerav",
	"Nishant",
	"Nitin",
	"Nivant",
	"Pranay",
	"Punit",
	"Pratap",
	"Prathamesh",
	"Pushkar",
	"Prathit",
	"Prahas",
	"Prajnan",
	"Prithvij",
	"Praruj",
	"Prathu",
	"Pruthviraj",
	"Pushpraj",
	"Rujul",
	"Sirak",
	"Sankalp",
	"Sanket",
	"Shantanu",
	"Shashank",
	"Siddhesh",
	"Sarvesh",
	"Subhash",
	"Sushil",
	"Swapnil",
	"Tushar",
	"Vasant",
	"Vijay",
	"Vishal",
	"Vivek",
	"Aaaqil",
	"Aafaaq",
	"Aaftab",
	"Aakif",
	"Aalim",
	"Aarif",
	"Adib",
	"Basir",
	"Baber",
	"Badeeh",
	"Baasim",
	"Baashir",
	"Chafik",
	"Daamin",
	"Daanish",
	"Dafiq",
	"Dahhak",
	"Diya Al Din",
	"Ehan",
	"Ebrahim",
	"Eijaz",
	"Ekhlaq",
	"Ekram",
	"Fadil",
	"Faaiz",
	"Faarooq",
	"Faeem",
	"Faateh",
	"Faakhir",
	"Gafur",
	"Ghaalib",
	"Ghaazi",
	"Gamal",
	"Hussain",
	"Haakim",
	"Haasim",
	"Habib",
	"Ijaz",
	"Jabeer",
	"Karim",
	"Liban",
	"Mastan",
	"Najeeb",
	"Owez",
	"Parvez",
	"Qadir",
	"Riyad",
	"Suhail",
	"Taabish",
	"Umraah",
	"Aakesh",
	"Ishaan",
	"Mayan",
	"Rayaan",
	"Vedant",
	"Aanan",
	"Charun",
	"Devaj",
	"Harsith",
	"Ishir",
	"Kanan",
	"Mahit",
	"Ojasvat",
	"Purnit",
	"Rodas",
	"Ekansh",
	"Hredhaan",
	"Jairaj",
	"Onkar",
	"Reyansh",
	"Samesh",
	"Viraj",
	"Yash",
	"Ahaan",
	"Ansh",
	"Arhaan",
	"Arin",
	"Arjun",
	"Arnav",
	"Aryaman",
	"Aryan",
	"Atharv",
	"Avi",
	"Avyaan",
	"Ayaan",
	"Ayush",
	"Ayushman",
	"Azaan",
	"Bachittar",
	"Bahadurjit",
	"Bakhshi",
	"Balendra",
	"Balhaar",
	"Balvan",
	"Balveer",
	"Banjeet",
	"Bhavin",
	"Brijesh",
	"Chakradev",
	"Chakradhar",
	"Champak",
	"Chanakya",
	"Chandran",
	"Chandresh",
	"Charan",
	"Chatresh",
	"Chatura",
	"Daksh",
	"Dakshesh",
	"Dalbir",
	"Darpan",
	"Darsh",
	"Darshit",
	"Dev",
	"Devansh",
	"Dhairya",
	"Dhanuk",
	"Dhruv",
	"Divit",
	"Divyansh",
	"Eeshan",
	"Ekalinga",
	"Ekapad",
	"Faiyaz",
	"Faraj",
	"Faras",
	"Farhan",
	"Fariq",
	"Faris",
	"Fitan",
	"Fiyaz",
	"Gatik",
	"Gaurang",
	"Gauransh",
	"Gaurav",
	"Gautam",
	"Girik",
	"Girindra",
	"Girish",
	"Gopal",
	"Guneet",
	"Harish",
	"Harsh",
	"Harshil",
	"Hemang",
	"Hitesh",
	"Hridaan",
	"Hriday",
	"Hritik",
	"Hunar",
	"Ikbal",
	"Imaran",
	"Isaac",
	"Ishwar",
	"Jagat",
	"Jagdish",
	"Jai",
	"Jason",
	"Jatin",
	"Jeet",
	"Jivin",
	"Kabir",
	"Karan",
	"Kiaan",
	"Krish",
	"Krishiv",
	"Krishna",
	"Laban",
	"Laksh",
	"Lakshay",
	"Lakshit",
	"Lavish",
	"Lohit",
	"Lucky",
	"Manan",
	"Manav",
	"Manbir",
	"Manthan",
	"Mohammed",
	"Moksh",
	"Naksh",
	"Nakul",
	"Naveen",
	"Neel",
	"Nihal",
	"Nimit",
	"Nirvaan",
	"Nitesh",
	"Om",
	"Omkaar",
	"Onveer",
	"Orinder",
	"Parth",
	"Parv",
	"Pranav",
	"Pranay",
	"Praneel",
	"Pranit",
	"Pratham",
	"Pratyush",
	"Qadim",
	"Qarin",
	"Qasim",
	"Rachit",
	"Raghav",
	"Ranbir",
	"Riaan",
	"Rishi",
	"Ritvik",
	"Rohan",
	"Romesh",
	"Ronith",
	"Ryan",
	"Saatvik",
	"Saihaj",
	"Saksham",
	"Samaksh",
	"Samar",
	"Samarth",
	"Sarthak",
	"Shaan",
	"Shaurya",
	"Siddharth",
	"Sohail",
	"Stuvan",
	"Suveer",
	"Taksh",
	"Tanay",
	"Tanish",
	"Tanmay",
	"Tarak",
	"Udarsh",
	"Ujjwal",
	"Utkarsh",
	"Vaibhav",
	"Veer",
	"Viaannew",
	"Vihaan",
	"Vivaan",
	"Wahab",
	"Warinder",
	"Wridesh",
	"Yagnesh",
	"Yatan",
	"Yatin",
	"Yuvraj",
	"Zaid",
	"Zain",
	"Zashil",
	"Zayan",
	"Zehaan",
	"Atharva",
	"Azad",
	"Baljiwan",
	"Bhaavik",
	"Chaitanya",
	"Chitaksh",
	"Divij",
	"Ehsaan",
	"Ekaraj",
	"Ekavir",
	"Ekbal",
	"Falan",
	"Faqid",
	"Frado",
	"Gagan",
	"Gunbir",
	"Hardik",
	"Himmat",
	"Ikshit",
	"Indrajit",
	"Ivaan",
	"Izaan",
	"Jainew",
	"Jaiyush",
	"Jowar",
	"Kalpit",
	"Kanav",
	"Kushagra",
	"Lauhit",
	"Maanas",
	"Maanav",
	"Madhav",
	"Medhansh",
	"Mitesh",
	"Nachiket",
	"Navneet",
	"Navodit",
	"Nishith",
	"Ohas",
	"Parthsarathy",
	"Purab",
	"Qabil",
	"Ranveer",
	"Rehaannew",
	"Rudranew",
	"Rudransh",
	"Rushil",
	"Sahil",
	"Sai",
	"Sathviknew",
	"Savar",
	"Shayak",
	"Shivansh",
	"Shlok",
	"Shray",
	"Tanveer",
	"Tanvik",
	"Teerth",
	"Tejas",
	"Udant",
	"Umang",
	"Upkaar",
	"Uthkarsh",
	"Virat",
	"Warjas",
	"Wazir",
	"Wriddhish",
	"Yayati",
	"Yug",
	"Yuvaan",
	"Zayyan",
	"Zeeshan",
}
