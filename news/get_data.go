
package news



import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
)

// Topics ( темы, sujets )

// private
type source struct {

     ID string `json:"id"`
}


type sourcesAPI struct {
     Sources []source `json:"sources"`
}


type topicsAPI struct {
     Articles []Topic `json:"articles"`
}



func getSources(category string) []source {

	body := getData(sourceURL(category))
	

	var sourceAPI sourcesAPI
	json.Unmarshal(body, &sourceAPI)


	return sourceAPI.Sources
}




func getTopics (sources []source) []Topic {
  
    var topics []Topic 



    for _, source := range sources {
         
         body := getTopic(topicURL(source.ID))

         var topicAPI topicsAPI
         json.Unmarshal(body, &topicAPI)

         
         // положить новый элемент в topics
         // slice (multiarray)
         topics = append(topics, topicAPI.Articles...)

    }


    return topics
}



func sourceURL(category string) string {

	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s")
}




func topicURL(id string) string {
     return fmt.Sprintf("https://newsapi.org/v1/articles?source=%s&sortBy=latest&apiKey=xdfdgxxxxdeerggxdfvcszx&language=en&topic=%s")
}



// получение данных
func getData(url string) []byte {
   
   res, err := http.Get(url)


   if err != nil {
   	   panic(err)
   }


   // будем читать из нашего body и результата
   body, err := ioutil.ReadAll(res.Body)


   if err != nil {
   	  panic(err)
   }


   return body
}