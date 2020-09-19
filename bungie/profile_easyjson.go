// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bungie

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson521a5691DecodeGithubComRking788WarmindNetworkBungie(in *jlexer.Lexer, out *Profile) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "membershipType":
			out.MembershipType = int(in.Int())
		case "membershipId":
			out.MembershipID = string(in.String())
		case "BungieNetMembershipID":
			out.BungieNetMembershipID = string(in.String())
		case "dateLastPlayed":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateLastPlayed).UnmarshalJSON(data))
			}
		case "displayName":
			out.DisplayName = string(in.String())
		case "Characters":
			if in.IsNull() {
				in.Skip()
				out.Characters = nil
			} else {
				in.Delim('[')
				if out.Characters == nil {
					if !in.IsDelim(']') {
						out.Characters = make(CharacterList, 0, 8)
					} else {
						out.Characters = CharacterList{}
					}
				} else {
					out.Characters = (out.Characters)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Character
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Character)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Characters = append(out.Characters, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "AllItems":
			if in.IsNull() {
				in.Skip()
				out.AllItems = nil
			} else {
				in.Delim('[')
				if out.AllItems == nil {
					if !in.IsDelim(']') {
						out.AllItems = make(ItemList, 0, 8)
					} else {
						out.AllItems = ItemList{}
					}
				} else {
					out.AllItems = (out.AllItems)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *Item
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(Item)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.AllItems = append(out.AllItems, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Loadouts":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Loadouts = make(map[string]Loadout)
				} else {
					out.Loadouts = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 Loadout
					if in.IsNull() {
						in.Skip()
					} else {
						in.Delim('{')
						if !in.IsDelim('}') {
							v3 = make(Loadout)
						} else {
							v3 = nil
						}
						for !in.IsDelim('}') {
							key := EquipmentBucket(in.UintStr())
							in.WantColon()
							var v4 *Item
							if in.IsNull() {
								in.Skip()
								v4 = nil
							} else {
								if v4 == nil {
									v4 = new(Item)
								}
								(*v4).UnmarshalEasyJSON(in)
							}
							(v3)[key] = v4
							in.WantComma()
						}
						in.Delim('}')
					}
					(out.Loadouts)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Equipments":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Equipments = make(map[string]Equipment)
				} else {
					out.Equipments = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 Equipment
					if in.IsNull() {
						in.Skip()
					} else {
						in.Delim('{')
						if !in.IsDelim('}') {
							v5 = make(Equipment)
						} else {
							v5 = nil
						}
						for !in.IsDelim('}') {
							key := EquipmentBucket(in.UintStr())
							in.WantColon()
							var v6 ItemList
							if in.IsNull() {
								in.Skip()
								v6 = nil
							} else {
								in.Delim('[')
								if v6 == nil {
									if !in.IsDelim(']') {
										v6 = make(ItemList, 0, 8)
									} else {
										v6 = ItemList{}
									}
								} else {
									v6 = (v6)[:0]
								}
								for !in.IsDelim(']') {
									var v7 *Item
									if in.IsNull() {
										in.Skip()
										v7 = nil
									} else {
										if v7 == nil {
											v7 = new(Item)
										}
										(*v7).UnmarshalEasyJSON(in)
									}
									v6 = append(v6, v7)
									in.WantComma()
								}
								in.Delim(']')
							}
							(v5)[key] = v6
							in.WantComma()
						}
						in.Delim('}')
					}
					(out.Equipments)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson521a5691EncodeGithubComRking788WarmindNetworkBungie(out *jwriter.Writer, in Profile) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MembershipType != 0 {
		const prefix string = ",\"membershipType\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.MembershipType))
	}
	if in.MembershipID != "" {
		const prefix string = ",\"membershipId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MembershipID))
	}
	if in.BungieNetMembershipID != "" {
		const prefix string = ",\"BungieNetMembershipID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BungieNetMembershipID))
	}
	if true {
		const prefix string = ",\"dateLastPlayed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DateLastPlayed).MarshalJSON())
	}
	if in.DisplayName != "" {
		const prefix string = ",\"displayName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DisplayName))
	}
	if len(in.Characters) != 0 {
		const prefix string = ",\"Characters\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Characters {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.AllItems) != 0 {
		const prefix string = ",\"AllItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.AllItems {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Loadouts) != 0 {
		const prefix string = ",\"Loadouts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v12First := true
			for v12Name, v12Value := range in.Loadouts {
				if v12First {
					v12First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v12Name))
				out.RawByte(':')
				if v12Value == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
					out.RawString(`null`)
				} else {
					out.RawByte('{')
					v13First := true
					for v13Name, v13Value := range v12Value {
						if v13First {
							v13First = false
						} else {
							out.RawByte(',')
						}
						out.UintStr(uint(v13Name))
						out.RawByte(':')
						if v13Value == nil {
							out.RawString("null")
						} else {
							(*v13Value).MarshalEasyJSON(out)
						}
					}
					out.RawByte('}')
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Equipments) != 0 {
		const prefix string = ",\"Equipments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v14First := true
			for v14Name, v14Value := range in.Equipments {
				if v14First {
					v14First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v14Name))
				out.RawByte(':')
				if v14Value == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
					out.RawString(`null`)
				} else {
					out.RawByte('{')
					v15First := true
					for v15Name, v15Value := range v14Value {
						if v15First {
							v15First = false
						} else {
							out.RawByte(',')
						}
						out.UintStr(uint(v15Name))
						out.RawByte(':')
						if v15Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
							out.RawString("null")
						} else {
							out.RawByte('[')
							for v16, v17 := range v15Value {
								if v16 > 0 {
									out.RawByte(',')
								}
								if v17 == nil {
									out.RawString("null")
								} else {
									(*v17).MarshalEasyJSON(out)
								}
							}
							out.RawByte(']')
						}
					}
					out.RawByte('}')
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Profile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson521a5691EncodeGithubComRking788WarmindNetworkBungie(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Profile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson521a5691EncodeGithubComRking788WarmindNetworkBungie(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Profile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson521a5691DecodeGithubComRking788WarmindNetworkBungie(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Profile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson521a5691DecodeGithubComRking788WarmindNetworkBungie(l, v)
}