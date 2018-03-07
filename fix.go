boolean := false
defer func() {
if rows != nil {
rows.Close()
}
}()
if rows.Next() {
boolean = true
}
return boolean, nil
