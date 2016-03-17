package server

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"strings"
	"time"
)

func ETagFor(data []byte) string {
	hash := sha1.Sum(data)
	b64hash := base64.StdEncoding.EncodeToString(hash[:])
	return "\"" + b64hash + "\""
}

type StaticHandler struct {
	Path     string
	MimeType string
	ETag     string
	Data     []byte
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" && r.Method != "GET" {
		w.Header().Set("Allow", "HEAD, GET")
		http.Error(w, "GET required", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set(ContentType, h.MimeType)
	w.Header().Set(CacheControl, CacheControlPublic)
	w.Header().Set(ETag, h.ETag)
	http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(h.Data))
}

var StaticHandlers []*StaticHandler

const cssStyle = ``

const jsMain = ``

const jsAutoclose = `window.addEventListener('load', window.close);
`

const icoFaviconB64 = `
AAABAAQAEBAAAAEACABoBQAARgAAACAgAAABAAgAqAgAAK4FAAAwMAAAAQAIAKgOAABWDgAAQEAA
AAEACAAoFgAA/hwAACgAAAAQAAAAIAAAAAEACAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAD54cgA
paWlAOPj/wD+/PkAKSlPAO+xcQD76dgA88KQAGho/gDe3NkA///+AOyeUADywIwA9MqfACQjRAD/
/fsAx8TCAPv59wD99u8A35BAACkpVwACAv8AMzP4ALq6/wBgYO0A8LV6APz8/ABvRx8ANCMSAPj4
+ADqmEUA9fX/AEpK/wDxuoIASUluANaJOwDo6OgAMzMzAOTk5ADi4uIAcHCCACgoMgDa2toAJSUl
ACMjIwALC+EAYF74AOHX6AC5uf8ALCzyAOuZRwAbG/8A+eLLAPG7gwD++PIAV1XiAOqXQwBLS/4A
KB/hAK6urgBFReQA+vr/AOXh9wDvrWoA/fbuAAAA/wD76tgACgryAPnjzAD//v4AKiglAFlZogB+
fn4A+vr5AOqYRAD54ckAMzP0ACAgVADx8PAA58emAP78+gDqlkEA9c6nACAVqABeXl4AXFxcAP76
9wDrm0kADAyeAP317ADy8v8AYGBxAO2lXADikUAA9MeZAP///wBMTEwAgoL/APn5+QDx8fEAEhL/
APC0dwDKyv8ADw+lAOqXQgD1z6gA+eDHAAEB+AAsLCwAKioqANvb2wAlJS8AlZX/AFBQqADtpl0A
dHT/AP7+/wDb2dgAx8fHANHPzgAcHFIA//79ADk5OgBVVWQAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAX18KD0VO
dxAJEQNWVlBQWQpfX3QdIk0pbFQqX19fX19CX19fZEFBQUxbenVFX19fDANfXxcfXwIgFXtVSV9Q
agYKX19fXxpudhYxLHZfAAsPX18KX2IoeG8EQRRINhlRUF9fX3k8QUFBGENYYEkNaERFX18wQWcn
WjlBLSVjaUpcBzRfYUErJl9fM0FGT3JRUVE/EnBBLAFfX3NBHBNRUVFROF4+QXFsOyQ3axtRUVFR
UVEeNS5BR20Oa1MjUVFRUVFRUVdLCEFBQTpdUVFRUVFRUVFRIUA9Zi8FUVFRUVFRUVFRUR5laVJl
MlFRUVFRUVFRUVFRUVFRUVFRUVFRUQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAIAAAAEAAAAABAAgAAAAAAAAEAAAAAAAA
AAAAAAABAAAAAAAAQD5cAAcA/wAXFoYAMTMyACwvLQArLSwAzKV8ANuNPgAaGBQAJykoABQB/wAj
JSQAlZO9AMiHPwAfISAAZEIZAM/S0ADzun8ALi//AB0fHgD63LwAGRsaAPz//gB5eJoA+teyAA4M
swAaFJcAgIL/AP369QBmZ2wA9cyjAPTy7ADWzOUA3+D/ABYWQwCtsK4ARUfAAGRg4wDqoVUACAHF
AJ+ioADyvokAJidXAJWYlgCRlJIA99SwAGtq+ADtmUEA/PTrAF5g/wBEQj0A+vLpAIyO/wD58OgA
MDDrACQjPgCNWSYA7qBJAJeZ/QDvplsA7aRZAOWRNgBbW7YA9sOOAE42HAAZEwsA9vf9APC9iAAU
EYwAAAb3AIiH/wDSysEA+/3/AAsFvAB+gP8A7ppDAC8xeQDy6/YA/O7iAIWCfQD96dkA+OreAKRp
MQD6/fsA6ppQAPj7+QAmFnoA+OPUAPb59wDIi3QApaX/ADY12wD1x5gAaWiJAPPFlgA4OjkA5unn
ADI0MwD38/UASj40ACosKwCGhv4AJignAPKtbQCZm7cA//jwAO2paAD1vYEAAAD/AEBBRQD44MQA
RUb/ACAkkwD23sIA8Ld8ABEKhwDIy8kA9de3AMbJxwD//fcA/vv2AMTHxQD8+fQA9dKtAPjPpgDr
lz4AvL+9AAAA6ADytHEAJCn8APCybwDusG0AHhQIAOzp5ACytbMA3N38ADQw/wAbGncAqKupAODd
2ABsa3YA1trYAPXBjAARE08AzNDOAA8IkAA9P3IA//74AE1LRgDBei8AjpGPAP/37gDsmEAAio2L
AP317AC+wsAAT09ZAICDgQAmIP8AoGUsAKZoYQCVmK4AGBXCAOyjWABqbWsAaGtpAGRnZQANDL4A
UzYUAJKPigBVV1YAwL/8AF1a/wCKjP4ASUjuAPnr3wApJvkA/OjYAF9e2QBaWJsAREdFAPTz/wDx
qF4APD89AOns6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAABYWFhYWkxx4k5NIeGI1TTU1YmIcYpeXmpqalxwcejBXFhYWFhYWFhZIQrh0iiup
qYp0NbVTFhYWFhYWFhYWFndIFhYWFhYWFhaFLJQDCQsJA5SMhoVVSBYWFhYWFhYWFncWFhYWFhaT
FmCzAhpwTDJfBQltqU1YFhYWFhYWFhYWUZNIFhYWFhasbGxsbGxsbDYXlGYDT4tiSBYWFhYWFkge
TpMWFhYWSG9sbGxsbGxsbGyypGRhqYWTFhYWSBYWk4MUl0gWFhYWZTGrtRYWh0YBbAGupgttI3hI
FhZIkxyvZxhikxYWFhYWFhYWFhYWFhaIbGyybQ6MixYWSEgzFD98r3eTFhYWFhYWFhYWSEhiQmKw
bGxdDl8jQkiTHHWAVK+TFhYWFhYWFhYWSGKFmyiplqFsbLBfC09gSEJQXlR9YhYWFhYWFhYWFhZi
kJ20BAsLA6JsAZIInBCTMHxqS32TFhYWFhYWFpMWeI2MiXNETG1tKmxsohNhimJOXiY9L2IWFhYW
FhYWSEgfW2xsbGxsbCWpRWwBFQmZYpctgy99mhYWFhYWFhaTeLBsAQFsbGxsbIGubGyPE4wfSLER
VH1Xd5NIFhYWFpM6bGwBomgcQloKbAFsbJETqotIUSk5fV5XYnhIkxYWUxJsAUltdkgWFpOIbGwB
SRWci2JxZ1QvtiktblFpSBYWbGxsIm12SBYWFkgBbAF/FZRHVymjSz1UVGdnKVeTFhYBbAETX4ZC
FhYWSK1sbAEVYwZnPJgvmD09Sz2jfE6TFkVsbGYLqbhIFhYWIWxsbEFADVRLL0s9S1Q9VEtyV5OT
iGwBABOUR2JISJNTbGwBQQ8HPVQ9S5g9PUsvS1SOUEIbbGwkZAsdeWJiSEJsbGyEOD09Sz0vPUtU
S1Q9VDtcUU1sbGwdZmaqLH55DGwBJ6iVPVQ9VC+YPT09PUs9SzuOFC5sbLClBQ5kMrenbGxWnz09
PZgvSz2YVEtUPVQ9VFRyV55sbAE+nAM3GWxsAVI9VEtUSz1LmD1LPUs9S5hLPTt1mp4BbGxsbGxs
bGygPUs9Sz1LPS89L5hUPVQ9VD2YJl5Ok0psAWwBbGywWVSYPVQ9VD0vL5hLPUs9Sz1LPS9UghSa
k0irNEY6IB5nL0s9Sz1LPS9LPT0vPVQ9VD1UPS9UKVczd3h4YpoUclSYPS89VD1UPUuYL5gvmC+Y
L5gvmEuja3tXV1dXexE8S5gvmC+YL5gvLz09Lz0vPS89Lz0vPUtUanJDQ3JqVC89Lz0vPS89Lz0v
fX19fX19fX19fX19fX2YmEsvmJh9fX19fX19fX19fX19fX19fX19fX19fX19fX19fX19fX19fX19
fX19fX19fX0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAwAAAAYAAAAAEACAAAAAAAAAkAAAAAAAAAAAAAAAEA
AAAAAABeYIgABwD/APj09gCcYiYAMDMxAOHk4gAuMS8A7Or+ACkrKgDbjT4AJiknAPro0AAUEdcA
ICMhAPW8gQA8P5gAOzxUAC0t/gBERf4A+ty8APz//gDHysgA6ppRAPbTrgC5tuIASEiDAIOA7gDz
8v8A+PbwAGRlagDzyqEAo6P+ALm8ugCUlpsA87VyABwU4wDvsW4A651FAB0d/wBiPnIA3d79AMiB
OQBOT1QAGxkWAGJgWwDV2dcA0dDdAAAD9gDv7f8APz9qACkr3QDtmUEAKyxPAD4/2wBOMg4AiYyK
ADcjPQD78uoAhYiGAPjw5wD87eEA+eveAD09/wDuoEkA8urhALO3tQB8fn0AY0WtACUe9AAOCqcA
dXh2AG9vegCzs/8AMzUzAAUF3gBDPf8Ai4j4ACEejAD7/f8AJiK2AGdo/wDumkMA//fvABocOAC9
qpsAlF4uANXU+gC9u7YACAjUABYXYwBWVb8AJR7/APr9+wBISkkA+eXVAPj7+QALBK8A8OLWAPCn
XQDt8e4A88WWAOrt6wAQAMUAycr/ACcg6gCwiJ8AGhP7ABQA/wBJSdQA76tqACIkIwA+J7gAUE/s
APrixgAAAP8AHiAfACsdDgDxuX0AEw1/ABYWKwAZHBoAuLr/ALiSaAD32bkA67N3APzo2QB4d5kA
//33APnWsQApJWoA/vv2AIODtQBpam8AdXiMAPv58wAlIiMA1tHUAMDDwQD4z6YA9Pj2AOuXPgDs
z64A6ZU8ALW21AAgI00A8u/qACUr/QBbSjwAGRpQALCzsQBWWFwAYUMkACAhJQCnqv8AGBO3ANXS
/wBPUNsAyrzjABgZHQDf3dcAoqWjADg3cADX1c8AHhvzAJmYpABXVVAAjY7AAJaZlwD//vgAFQqn
AOyYQADUhzUAlJT9AEVEzgCAUSYA/fXsABEUewA7PIsAgoWDAB4VCQA5NvsA5uTfAAoA5gDGyeAA
5pI3AOyjWACqp6IAcHNxAAgDxgAaFAwA9/j+APPAiwBqbWsAFRu8AJKP/wB5Ty0AFxc4AHl3/QBk
Z2UAbG3LACwuLABwcP4Ah4b+AGRklADj5v4AwnsxAENE1wDj5OoADRKWAFNQ5AAFAO8Ac3b+AP3v
4wA7OoIATE9NALh3OACGh6MAQ0VEAGJgtADv8vAA9siZAAAA5wBpaMwA1tLuADk7OgA2OTcAQj7k
ACYkIADl6OYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAFBQUFBQUFKioAgICqL6oqE6GUlKvQD091NQcUlJSUgKvOTk7Ozs5UgICAlIcO317FBQUFBQU
FBQUqKioTlwCvgICz4iIiIiin88CqKgUqKioqKioqKioqKioqBSoqKg5FBQUFBQUFBQUFBQUFKi+
zxWVITpCQkJCIaRXiM+GAk4UFBQUFBQUFBQUFBQUFBS+FBQUFBQUFBQUFBQUFFyRV0Kl2eHICMgE
4CosQrouY1wUFBQUFBQUFBQUFBQUFBSoThQUFBQUFBQUFBQUFKiIN9njc+PjDW6Y44cI4JY6V88C
ThROFBQUFBQUFBQUFBQUUqgUFBQUFBQUFBSoqF+3WsFYtkoMT7GlKuAIhwjghLothk6oFBQUFBQU
FBQUFBQUPYaoFBQUFBQUFBSoeRFycnJycnJycnIBNX7AXQaHSSy6iAKoTk4UFBQUFBQUFBROgNQC
qBQUFBQUFBQUPnJycnJycnJycnJyci/R2CzgmEksugWGTqgUFBQUFBQUqBSoZBPUqBQUFBQUFBQU
wnJycgFycnJyAQFycnJyaqa74IcEhFeRvqgUFBQUFBROTr6CbR5eUqgUFBQUFBQUVnImyXkHFBQb
Z9NbcnJyci+mLEnj2TqIixQUFBQUTk6ohlI8Yr8T1AIUFBQUFBQUG0gUFBQUFBSoFBQUHxEBcnJr
g6UIhyq65AIUFBQUqE5SPF4XbdxxO38UFBQUFBQUFBQUFBQUFBROThROqMyScnJyJn7gDUlHotsU
FBROvlJhe2R1v3t9AqgUFBQUFBQUFBQUFBQUqBQUFKhci19fJnJycs4smIcqus8CFBROhjx7v225
E9SGqBQUFBQUFBQUFBQUFBRcvk4CAs+In4iIGAFycnLLSZhJQqKLFE5OAl5kbRaqPQKoFBQUFBQU
FBQUFBQUFBQUvgKIQSFCQkdCN2xycnJoXePjLJUCqBSCPRcOuaqMr6gUFBQUFBQUFBQUFBQUqKgC
kSA6LNngyMgE4DHScnIBGZiH2ae1vgLUE79iUYyMghQUFBQUFBQUFBQUFKgUFE7PQbvZCg0r4+Nu
4+O8cnJyDAjjyEYu21ILHiQWjIyMAhQUFBQUFBQUFBQUTk6of2UghE3BWEpYmtWlKtmBcnJy0pAr
CB1XAn99FyQljIyMgqgUFBQUFBQUFBQUFE4Uiy6tAXJycnJycnIvxyFGL3JycrCehyq6Y6jUew65
FoyMUqgUFBQUFBQUFBQUTqh/30RycnJycnJycnJycrSPznJyAWCeh+Cntb6+PR4kFoyMPAIUFBQU
FBQUFBQUFKiCW3JycnJycnJycnJycnLSGnJyckp3K+CyLr6o1IAOPzOMXlKoThQUFBQUFBQUqKjT
cnJycnK22i5fB0gSAXJyAXJycnLEh8hCooZOO4AiFlGMjX1SAk6oThQUFBQUFMxrcnJyAUWlurWo
FBQUeRFycnJycnJZKwjGiAICC4okFqqMDopeOVICf6ioqBQUFHkBcnJy3YcqQRtcFBQUFDBqcnJy
cnLQeIcdiQI5XmRtFoyMbQ4ee159Pa8CThQUFKwBcnJyduMqV9tOFBQUFBRWcnJycnJmK4cdV2Fe
iiIWjI6qURYkv78eF147ThQUTspycnJyU4eloJGoFBQUFE6+ynJycnJKK4cqVI0ebRYWqhaOjhYW
FhZiDh59gqgUFMpycnJy4yvZIbVfThQUFBRfzGtycnIvvYeTenxtFhYWjI6qFo4WFhYWFg6A1KgU
qKxycnIBCJhJQi6GFBQUFBROvhJycnIBvbOX1xYWFo4WjhaOjoyOjI4zFm1kfQKoFJlycnJyEIfj
lrplvhQUFBQUqFBycnIvs3TDzRaOjI6MjqqMFo4WjhaOFhYOezy+qJtrcnIBDwiHSUJXHBROFE4U
X8VycnIBsziuzbgWjhaOFo6MjoyOjI6MjlFiZBPUTk4mcnJyMtmHDV06iM++f4JfHMVycnLSszYD
qxaOjI6MjqqMFo4WjhaOFo4WbR5ePE7FcnJyLwDgh4cqQleIzwICz3BycgFmdMPXq44WjhaOFo6M
jo6MjoyOjI4WFiTcez05AXJycuIdBIcK4B03oJWVpCNycnJ2NlUpFoyOjI6MjqqMjBaOFo4WjhaO
FhZtZBd9THJycnLeLOCH4wjhXaWWoXJyAXI4rs24FrgWjhaOFo6Mjo6MjoyOjI6MjjNRYiQXYUty
cnJynIXWyJiHh+OU3XJyAWauzasWjqqOjI6MjqqMjBaOFo4WjhaOFo4WjG0OE98vcnJycqOtAKXg
NKkvcnJycifNCRaOFo4WjhaOFo6Mjo6MjqqOjI6MjoyOjhZtFwtWawFycnJycnJycnJycnIBb6u4
Fo6MjoyOjo6MjqqMjBaOFo4WjhaOFo4WjBZtv15SKBIBcnJycnJycnIBcnJDuI4WjhaOFo4WjBaO
Fo6Mjo6MjoyOjo6MjoyOjI65DhPUAr5IJnIBAXJyAQFyamm5Fo6Mjo6OjI6Mjo6OjI6qjBaOFo4W
jBaOFo4WjBZRJB5eUqgUFGfTPluSS8mdFw4WURaOFowWjhaOFowWjhaOFo6OjI6qjo6OjI6qjo6O
uXWNPYYCThROFKioAjte3G0WFoyOqo6OjoyOqo6OjqqMjowWjBaOFowWjBaOFowWFm2/gH2vhgKo
hr4COV4Xv7lRFowWjhaMFowWjhaMFo6Mjha4jo64jo64jo64jo64jlFtdR4TfT3UPT1ee4p1bRYW
uI6OuI6OuI6OuI6OuI6MjIwWFo4WFo4WFo4WFo4WFhYWbSS/ZIqAF4rcvyRtFqqMFhaOFhaOFhaO
FhaOFhaOjIyMjIyMjIyMjIyMjIyMjIyqURZtJCQiJCRtbbkWqoyMjIyMjIyMjIyMjIyMjIyMjIyM
jIyMjIyMjIyMjIyMjIyMjKqqURYWFhYWqqqqjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyM
jIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyM
jIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAKAAAAEAAAACAAAAAAQAIAAAAAAAAEAAAAAAAAAAAAAAAAQAAAAAAAAYA/gBM
SP8AZ0MVAN/i4AASF/0AmJaRANuNPgAlJyYA+ujQAMfH+gDT1tQA7qppACAjIQD748cA9LyAAD8n
KQDwuHwA+ty8ABgbGQD8//4Aw8bEAEVHigD+9ewAZmdsAL/CwABPUOgAGRsrAF5f/gCUlpsAHRRe
AGlKLADyxZYAGhDhAO6xbQCqq8kARUZ9AB4bGQDGfzcAp6qoAE5LjQD2wo0AEwD/APn6/wDa2NIA
WVtcAFdW6gCKidEA5KpqANaTVQCuqaIANzdRAD8/twAzIN0A7ZlBACUm4wAJALkAxcK9AH9QJQD7
8uoA/u/jAL+8twD48OcATk5YAIOGhAD6698Atrb7ABUUYQBfXp8A7qBJACAWCwDn5/QAVFH/AGdb
UAChZi0AvLXjAOC+mQBpZ2MA3d3qAO2kWQBzdnQANiOrAGtrwwA4NsgAX11ZADc5NwCwr/wA8r+K
AAEK+AAsL0AAWFd3AF88TwCGhf0ANSYXAFteXADumkMA8q9iAH18/gBPUlAAgH+WAKur4AAREasA
SUxKAPr9+wDqmlAADQyxAPnl1QD4+/kAnZ6jAENGRADThTIAMzEtAAIB6QCdfLcALy0pAPXHmAAP
EHcALSsnAOzv7QAbFg4A5+voADk6+QAzNjQA5OflABwZ/QAvMjAAnZuWAPXx8wBOTF0Ar24sACsu
LADa3dsAHxzOAPKtbQDP0O4AfX6DACIkIwAAAP8AHiAfAPjgxACHhYAAKi6uAHRMHADIy8kA//33
APjUsAD9+/UA+/nzAPjPpgDrlz4AioyKAK2IXQALCZwA9vPuAPD08gA7O9cAuLu5APK0cQAoKnAA
KCktAO/t5wAEB84A2c/oAOyeRgDd3/0A6OXgAAoSxwAwLq4A4N3YAI6L8wAwMd0Ad3ScAKCjoQA7
Pu8AGBWPADM2+wBJMBYAIB8iAM7S0ACwmn8AREX/ANDNyACbnssANCkfAPXYuAD//vgAlZb+ADw9
4wAzIQoAwsD9APHw/gBnUj8AFhgzAAAA8wAiIDwAsHpDAEZIRgB6fXsA8KdcAOaSNwDqoVYAGgew
APP0+gApLIoAQkHPAC8vOQB6e7EAuHhZAF9hYABvbv0AZGSUACgqKAD/+O8A2tn/ALy8yQBAPk4A
BwDxAFVU0gBzdv4Akl0sAIqM/gBwbX8AJxkJAAUA2wCMjakAMS1dAP3o2QC4trEAJR//AHV22AC2
dTYA9fj2AOqVPABAQ0EA9MuhAOno/gAVE+8AOz08ACIl+QDh08cAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATExMTExMT
ExO4Kip+0yp+Ko+Pj7iRkZKROj07pDo6PT0709M6fjrTOtMWOxY7Ozs7Ozs7OtOSfpJ+0xY7CGmQ
ExMTExMTExMTE7i4uLi4uLgTKioqE5IqfnWffqSkpKSfn35+j4+PuLi4KioqKioqKioqKo+4uLi4
uCoqKio6QBMTExMTExMTExMTExMTExMTKriRmXWkTbQ4ODziPDw8OLRNpJnm5hMqExO4uLi4uLi4
uLi4ExMTExMTuLi4uDoTExMTExMTExMTExMTExMTKrgqmaS01X0FhoYXFxdPTIscfZuOp5+YkbgT
uBMTExMTExMTExMTExMTExMTEyq4ExMTExMTExMTExMTExMTExO4agPVHIZTYcPsVG5ublTDPlMX
izHVTcnmKioqExMTExMTExMTExMTExMTExMTuLgTExMTExMTExMTExMTExMqE34rMUzsbp6wB7Cw
iYewsHRuVD5TizGORpkTExMTExMTExMTExMTExMTExMTExPTuBMTExMTExMTExMTExMTE2Z+sQVh
bgyHB9LSgW6ennR0sLCebj4XHDinmCq4ExMTExMTExMTExMTExMTExMTkiq4ExMTExMTExMTExMT
Kiq45mPYg97XV1fX64OMFVNhw+xunnR07CyL4k2fKrgqKhMTExMTExMTExMTExMTE0DTuCoqExMT
ExMTExMTE7ij2QSIiIiIiIiIiIiIiIgEmqqGU8NusJ5uU4bip8mRKri4ExMTExMTExMTExMTExMR
QNO4uBMTExMTExMTExMTswCIiIiIiIiIiIiIiIiIiIiIut/EU+x0B25hi+Knfri4KhMTExMTExMT
ExMTE7gT6YpAKioTExMTExMTExMTE2CIiIiIiIiIiIiIiIiIiIiIiIjA5BxMw56wbj6VOKSSKiq4
ExMTExMTExMTuCoqKhCQaTvTuBMTExMTExMTExNBiIiIAIiIiIiIiIiIAIiIiIiIiIgZaxfDdLDs
XRyOyZG4ExMTExMTExMTEyq4uJEhKJAI0yoqExMTExMTExMT6gCIVwHbQaO9Kr3UvFuuV4iIiIiI
iKxrTOx0dOxMMafmKhMTExMTExO4uCoqapI9ZxDpDT2PuBMTExMTExMTExOzub0TExMTExMTExMT
E9RHiIiIiIiIrBxTzIeePovVpI+4ExMTExMTKio6OjvhEccQHxE70yoTExMTExMTExMTExMTExMT
ExMTExMTExMTE7wEiIiIiIgZPz5xsG5TayvJKioTExMTKriRPUCKkB8hKJBpO5G4ExMTExMTExMT
ExMTExMTEyoTExMTuLgTE7gT1O2IiIiIAFFMVIewbIabn48TExMTuBN+OwiQH1acKHIR4dMqKhMT
ExMTExMTExMTExMTExMTuBMqExMqKip+an7qBIiIiIhXYsPSsG5TqyuZKhMTEyoqOkC3HxBfx5Np
QJKPuLgTExMTExMTExMTExMTExMTExNmarh+mXWfd34DRrxXiIiIiLoXzImebIbVdWoTKiq40z0R
6ZzHZ14NOzqPuBMTExMTExMTExMTExMTExMTE7i4Kip+Tae0ODw84jw4LQCIiIiIQ+y2iW5MMYLJ
uBO4kphpk1bFZzWUCCq4ExMTExMTExMTExMTExMTExMTExMqKn6njuIchoYXFxeGT2LAiIiIiDZT
nomeZQW0mCoqatNAtx+cx2fnlDoquBMTExMTExMTExMTExMTExMTE7i4fsmx4gVMU8PsVG58VOw+
pgCIiIiIJ26HsOyGm0aPuJJAaZMQC2dn55TTuBMTExMTExMTExMTExMTExMTEyoqKp+0McQ+7HGH
iSSJsLB0dFgAiIiIiIPsiYluTGukfioW4ZBWhGc15+eU0yoTExMTExMTExMTExMTExMTExMTkZmx
Jk/DnocSsIeenm5ucZ6eaIiIiIgAMrCwdD4FK3WRPYqQEE6i52fnlJK4ExMTExMTExMTExMTExMT
EyoqEyqn4tzKaKBvAMBvpYwjU11hw52IiIiIiK2JEofshhR+KtNAkCgLZ5Tn55TTuBMTExMTExMT
ExMTExMTKri4uJF11ctXiIiIiIiIiIiIiO1RBYYXb4iIiIigJLCwzE+bpI/TOxEfIceilOeUkioT
ExMTExMTExMTExMTKri4KiqSSuuIiIiIiIiIiIiIiIiIiK4iJqmIiIiIiL8SJG5Ta0bmKtNAk1Zf
Z2fnlDqPuBMTExMTExMTExMTExO4uCoqoQCIiIiIiIiIiIiIiIiIiIiIKWPkiIiIiIhzEiQHU2un
5ri4PZAfhGeU55RA07gTExMTExMTExMTExMTExO4RnuIiIiIiIgAAIiIiIiIiIiIiIgAqIiIiIiI
aCQknj6VsZgqKjtpH5zHZ+eUaTsqKrgTExMTExMTExMTExMTuNCIiIiIiIgA1zO1TcmjvNB7iACI
iAAAiIiIiG8SErDDixR+Ko89ipOcxzXnlJBAOo8qKhMqKhMTExMTExMTuOoAiIiIiACIaD4cCslm
ExMTKltXiIiIiIiIiIiIvySw7IY4d7iRO2konMdn55Qft+E7kY+PKri4KhMqExMTExO5iIiIiIiI
N55da6eSKhMTExMTCQSIiIiIiIiIiEIksOxPPMkqfkC3KC9ElOeUECiQaeE6kjoqkbi4uBMTExO4
R4iIiIgAwBpxFzF6KrgTExMTExPUKYiIiIiIiIitErDsFzykFj1p6RDFZ2fnlMWcH5C3aWlAPTrT
KioqExMTE+OIiIiIiJewbhcxTWYTExMTExMTE1WIiIiIiIiIZCSwbhcxp0Bpkw6ETmdn55SixZwQ
H5OQkBENPZi4uBMTExMEiIiIiIhCJJ5Ta6R+ExMTExMTExMTG4iIiIiIAKUSJHxMMe6Kk1YhTmc1
5+eUZ2fHxSGcnBAfkGk7j7gTExMTAIiIiIiIwQywUxwKmRMTExMTExMTE9RXiIiIiIjeJCR0SLJL
HxCcxWdn5+fnlOdnNWdnxcVfhCiQQDoquBMTuACIiIiIiJ6JnsOVtMkqExMTExMTExO4s4iIiIiI
byQktr6WLy/FRGdn5+dn55Tn5+dnZ5RnZ2cv6Yo7KrgTExMpiIiIiIjMJLDsTzx3argTExMTExMT
E1uIiIiIiNd2RVwewjDHZ2dn5+fn5+eU52fn5+dn55SiCxCQCNO4uBO4eIiIiIgAMp4kblMcK5gT
ExMTExMTE7gJiIiIiIiIdkW7jeUGZ5Tn52fn52fnlGfn5+fn5+fnZ2eEHxFAjyoquBuIiIiIiCNx
JJ7shjxN5ioqExO4ExMqo4iIiIiIiEXdrznOBmfnZ+fn5+fn5+fn5+fnZ+fn5+eixRCQaT2PuCq5
iIiIiIim7LCwbiyVFH6YuBMTKhMTuCqIiIiIiNd2Ra/aJQbn5+fnZ+dn5+eU5+fnZ+fn52fnlGcv
KJBAmCoqo4iIiIiINixusIfsTGu0dyoqahMTKhPJiIiIiIjeRd0CSW0G5+dn5+fn5+fnlOfnZ+fn
52fn52eURJwfEUDTj7iuiIiIiIjRw54kdOxMHNWnd5iRj+Z+hYiIiIiIoEWvjYAG5+fn5+dn5+fn
Z5TnZ+fn52fn5+fn52cLIR8R4Toq24iIiIhXrEzodLCw7FOLMRgrTQNNKy6IiIiIiJe7AtolBmfn
52fn5+fnZ+fn5+fn5+fn5+fnZ+eUogsQH5BpPTopiIiIiABRTOyHDAdUPkyLHDFrMX1SiIiIiIgd
rzmABufnZ+fn52fn5+fnlGfn5+dn5+fnZ+fn52dnxZxWkIpAqIiIiIiIV80Xw3SwsAdU6D4sF1N/
V4iIiIiID41JJcbnZ+fnZ+fn5+dn55Tn52fn5+dn5+fn5+fnlGdOhBCTEUB4iIiIiIgEzRc+bnSw
sHTSdNJ0aIiIiIiIyALaJcZn5+fn5+fnZ+fn5+eU5+fn52fn5+dn5+dn5+eUZ6ILEHKKoSmIiIiI
iFfYhs9seXQMDCSwZIiIiIiIAFpJJcYGZ+fn5+dn5+fn52fnlOfnZ+fn52fn5+dn5+dn5+dnogso
kAhBV4iIiIiIiATLQ1k+1uCtb4iIiIgAAFBJJQZn5+fnZ+fn5+dn5+fn55Tn5+fn5+fn5+fn5+fn
5+dn52fHnOlpOlUAiIiIiIiIiIiIiIiIiIiIiIiIACDOJcZn5+fn5+fn52fn5+fnZ+fn52fn5+dn
5+fnZ+fn52fn55SUZ4QftzuRvHuIiIiIiIiIiIiIiIiIiIiIiN7OBudn5+fn52fn52fn52fn5+fn
lGfn5+dn5+fnZ+fn52fn5+fnXqILKJBAkiqjs4iIiIiIiIiIiIiIiIgAiDQwZ2fn5+fn52fn52fn
5+fn5+dn55Tn5+dn5+fnZ+fn52fn5+dn58Y1x5yTijuRuBO8swCIAACIiAAAAACI63Ahx16U52fn
52fn5+fn5+dn5+fn5+eU5+dn5+fnZ+fn52fn5+dn5+fnZ8dfH5BAkrgqExPU20ftAAAp43hgSrcf
nGdnZ+fn52fn5+dn5+dn5+fn52fnlOfn5+fn5+fn5+fn5+fn5+fn55Q1xxDpETvTuLgTExMTE7gq
KiqSPWmTDl/HlOdn5+fn5+fn5+fn5+fnZ+fn55TnZ+fn52fn5+dn5+fnZ+fn52fnZ8cLEOmKQJiR
KioqKrgqKriRPUC3HxBOZ5Rn5+dn5+fnZ+fn52fnZ+fnZ+eU5+fnZ+fn52fn5+dn5+fnZ+fn5+c1
TpwokIrhO9ORKtMqkdMWQOERkxBfxzVn5+fn5+dn5+fnZ+fn5+fn5+fnlOdn5+fnZ+fn52fn5+dn
5+fnZ+fnNWfHX1aTkGlAQD09PUDhaWm36Q6ExWdnZ+fn52fn5+dn5+fnZ+fn5+dn55Tn5+dn5+fn
Z+fn52fn5+dn5+fnZ+eUZ05fnB+TkBERDRGKt7fpVpxfx2dn5+fnZ+fn52fn5+dn5+fnZ+fn5+eU
5+fn5+fn5+fn5+fn5+fn5+fn5+fn52deZ8ecnA5WHx8fHyhWnJwLx2dn5+fn5+fn5+fn5+fn5+fn
5+fn52fnlOdn52fnZ+dn52fnZ+dn52fnZ+dn52fnZ2dnRMULISGcnJxfxcfHZ2dn52fnZ+dn52fn
Z+dn52fnZ+dn52fn55Tn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fnNWdnZ2dnZ2dnZ2dnZzU15+fn
5+fn5+fn5+fn5+fn5+fn5+fn5+eU5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+eUlJSUlJTn
5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fnlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSU
lJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSU
lJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUlJSUAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=
`

func removeWhitespace(r rune) rune {
	switch r {
	case ' ', '\t', '\r', '\n':
		return -1
	default:
		return r
	}
}

func init() {
	icoFavicon, err := base64.StdEncoding.DecodeString(strings.Map(removeWhitespace, icoFaviconB64))
	Must(err)
	StaticHandlers = []*StaticHandler{
		&StaticHandler{
			Path:     "/css/style.css",
			MimeType: MediaTypeCSS,
			ETag:     ETagFor([]byte(cssStyle)),
			Data:     []byte(cssStyle),
		},
		&StaticHandler{
			Path:     "/js/main.js",
			MimeType: MediaTypeJS,
			ETag:     ETagFor([]byte(jsMain)),
			Data:     []byte(jsMain),
		},
		&StaticHandler{
			Path:     "/js/autoclose.js",
			MimeType: MediaTypeJS,
			ETag:     ETagFor([]byte(jsAutoclose)),
			Data:     []byte(jsAutoclose),
		},
		&StaticHandler{
			Path:     "/favicon.ico",
			MimeType: MediaTypeICO,
			ETag:     ETagFor(icoFavicon),
			Data:     icoFavicon,
		},
	}
}