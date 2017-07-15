(* learn the language by going through all the tutorials on ocaml.org *)

open Core

let x = 11

let square x = x * x

let sum_if_true test first second =
  (if test first then first else 0)
  + (if test second then second else 0)

let first_if_true test x y =
  if test x then x else y

let me = ("alex", 31)

let languages = ["javascript"; "golang"; "ocaml"]

(* method to remove sequential duplicates from a list *)
let rec destutter l =
  match l with
  | [] ->  []
  | hd :: [] -> [hd]
  | hd1 :: hd2 :: tl ->
      if hd1 = hd2
      then destutter (hd2 :: tl)
      else hd1 :: (destutter (hd2 :: tl))

let log_entry maybe_time message =
  let time =
    match maybe_time with
    | Some x -> x
    | None -> Time.now ()
  in
  Time.to_string time ^ " -- " ^ message

(* record types *)
type point2d = { x : float; y : float }

let magnitude { x; y } =
  sqrt (x ** 2. +. y ** 2.)

let distance v1 v2 =
  magnitude { x = v1.x -. v2.x; y = v1.y -. v2.y }

type circle_desc = { center: point2d; radius: float }
type rect_desc = { lower_left: point2d; upper_right: point2d }
type segment_desc = { endpoint1: point2d; endpoint2: point2d }

(* variant types *)
type scene_element =
  | Circle of circle_desc
  | Rect of rect_desc
  | Segment of segment_desc

(* test whether a point is inside a list of scene_ elements *)
let is_inside_scene_element point scene_element =
  match scene_element with
  | Circle {center; radius} ->
      (distance center point) < radius
  | Rect { lower_left; upper_right } ->
      point.x > lower_left.x && point.y > lower_left.y && point.x < upper_right.x && point.y < upper_right.y
  | Segment _ -> false

let is_inside_scene point scene_elements =
  List.exists scene_elements
    ~f:(fun el -> is_inside_scene_element point el)
(*
let rec read_and_accumulate accum =
  let line = In_channel.input_line In_channel.stdin in
  match line with
  | None -> accum
  | Some x -> read_and_accumulate (accum +. Float.of_string x)

let () =
  printf "Total: %F\n" (read_and_accumulate 0.)
*)
let languages = "Ocaml,Perl,C++,C"

let dashed_languages =
  let languages = String.split languages ~on:',' in
  String.concat ~sep:"-" languages

let (ints, strings) = List.unzip [(1, "one"); (2, "two"); (3, "three")]

let (+!) (x1, y1) (x2, y2) = (x1+x2, y1+y2)

let concat ?(sep="") x y = x ^ sep ^ y

exception Problem of int

let f x y =
  if y = 0
  then raise (Problem x)
  else x / y

let g x y =
  try f x y with Problem p -> p
