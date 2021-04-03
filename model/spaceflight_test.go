package model

import (
	"fmt"
	"testing"
)

func TestAbc(t *testing.T) {

	data := []byte(`
	[{"id":"6051e86631c42cd69c01e29a","title":"Russia looks to China as new space exploration partner","url":"https://spaceflightnow.com/2021/03/15/russia-looks-to-china-as-new-space-exploration-partner/","imageUrl":"https://mk0spaceflightnoa02a.kinstacdn.com/wp-content/uploads/2021/03/50875730681_b4e2d8c6cc_k.jpg","newsSite":"Spaceflight Now","summary":"Russia’s decision to partner with China on a planned lunar research station, and not join the U.S.-led Artemis moon program, was disappointing after more than two decades of cooperation on the International Space Station, says NASA’s top human spaceflight official.","publishedAt":"2021-03-15T10:25:18.000Z","updatedAt":"2021-03-17T14:07:25.840Z","featured":false,"launches":[],"events":[{"id":"108","provider":"Launch Library 2"}]},{"id":"6051e86631c42cd69c01e29b","title":"SpaceX extends its own rocket reuse record on Starlink launch","url":"https://spaceflightnow.com/2021/03/14/spacex-extends-its-own-rocket-reuse-record-on-starlink-launch/","imageUrl":"https://mk0spaceflightnoa02a.kinstacdn.com/wp-content/uploads/2021/03/l21_1.jpg","newsSite":"Spaceflight Now","summary":"Using a Falcon 9 booster flying for a record ninth time, SpaceX’s swift sequence of launches continued Sunday with a predawn liftoff from the Kennedy Space Center carrying another set of 60 Starlink internet satellites, the company’s third mission in 10 days.","publishedAt":"2021-03-14T10:25:18.000Z","updatedAt":"2021-03-17T13:18:13.974Z","featured":false,"launches":[],"events":[]}]
	`)

	result, err := UnmarshalJSON(data)
	fmt.Printf("After unmarshal: %#v\n", result)
	if err != nil {
		t.Error("Failed converting to array of structs.")
	} else {
		if len(result) != 2 {
			t.Errorf("got %v", result)
		} else {
			fmt.Println("result ", result)
		}
	}

}
