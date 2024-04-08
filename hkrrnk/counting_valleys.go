/*
An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly  steps, for every step it was noted if it was an uphill, , or a downhill,  step. Hikes always start and end at sea level, and each step up or down represents a  unit change in altitude. We define the following terms:

A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.
Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.

Example

 

The hiker first enters a valley  units deep. Then they climb out and up onto a mountain  units high. Finally, the hiker returns to sea level and ends the hike.

Function Description

Complete the countingValleys function in the editor below.

countingValleys has the following parameter(s):

int steps: the number of steps on the hike
string path: a string describing the path
Returns

int: the number of valleys traversed
Input Format

The first line contains an integer , the number of steps in the hike.
The second line contains a single string , of  characters that describe the path.

Constraints

Sample Input

8
UDDDUDUU
Sample Output

1
Explanation

If we represent _ as sea level, a step up as /, and a step down as \, the hike can be drawn as:

_/\      _
   \    /
    \/\/
The hiker enters and leaves one valley.

*/


// Logic is when taking an uphil step makes the distance == 0, that means hiker just came out of valley 
// Initialize two integer variables: vcount to keep track of the number of valleys traversed, and distance to keep track of the current altitude relative to sea level.
//Iterate through the characters in the path string.
//Increment distance for each uphill step ('U'), and decrement it for each downhill step ('D').
//If distance becomes 0 after an uphill step, it indicates the hiker has just finished traversing a valley. //Increment vcount.
//Return the final vcount as the result.
func countingValleys(steps int32, path string) int32 {
    // Write your code here
    var  vcount,distance int32
    vcount = 0
    distance = 0
    
  for i:=0;i<len(path);i++{
        if string(path[i]) == "U"{
            distance++
            if (distance == 0) {                    
                vcount ++
            }
        }
        if string(path[i]) == "D" {
            distance--
        }        
  }
    return vcount
}
