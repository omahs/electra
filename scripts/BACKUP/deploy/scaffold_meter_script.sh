ignite scaffold module meter --dep bank -y
ignite scaffold map meterreadings phase:uint whin:uint whout:uint mvolt:uint mhertz:uint mpf:uint maxmi:int --index deviceID:string,timestamp:uint --module meter --no-message -y
ignite scaffold message record phase:int whin:uint whout:uint mvolt:uint mhertz:uint mpf:uint maxmi:int --desc "Post a meter reading in the chain" --module meter -y
ignite scaffold map meterdirectory model:string installationdate:uint address:string ownerlastname:string  ownerfirstname:string ownerphone:string gpsjson:string active:bool triphased:bool monophasenb:uint --index deviceID:string,barcodeserial:string --module meter --no-message -y
ignite scaffold map contracts startdate:string enddate:string isdefaultperid:bool contractname:string --index fromDeviceID:string,toDeviceID:string,isactive:bool --module meter --no-message -y
ignite scaffold query listrecordings deviceID:string start:uint end:uint byblock:bool --response listrecording,total:uint --module meter --desc "List the queries from START and to END  when BYBLOCK is tru then filtering start-end parameters are interpreted as blocks otherwise assoonas BYBLOCKK equals fals consequently params get seen beeing DateTime timestams" -y