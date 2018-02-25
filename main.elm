module Main exposing (..)

import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)


main =
    Html.beginnerProgram { model = model, view = view, update = update }



-- MODEL


type alias Duration =
    { days : Int }


type alias Task =
    { name : String
    , duration : Duration
    }


type alias Model =
    { scratch : Task
    , tasks : List Task
    }


model : Model
model =
    Model
        { name = "", duration = { days = 0 } }
        []



-- UPDATE


type Msg
    = SetTaskName String
    | SetTaskDuration String
    | AddTask Task
    | RemoveTask Task


update : Msg -> Model -> Model
update msg model =
    case msg of
        SetTaskName name ->
            { model
                | scratch = { name = name, duration = model.scratch.duration }
            }

        -- TODO throw an error on duplicate task name
        SetTaskDuration duration ->
            { model
                | scratch =
                    { name = model.scratch.name
                    , duration =
                        { days =
                            String.toInt duration |> Result.toMaybe |> Maybe.withDefault 0
                        }
                    }
            }

        -- TODO throw an error on empty task name
        -- TODO throw an error on duplicate task name
        -- TODO throw an error on invalid task duration
        AddTask task ->
            { model | tasks = List.append model.tasks [ task ] }

        RemoveTask task ->
            { model | tasks = List.filter (\n -> n /= task) model.tasks }



-- VIEW


view : Model -> Html Msg
view model =
    div []
        [ input [ type_ "text", placeholder "Name", onInput SetTaskName ] []
        , input [ type_ "number", placeholder "Days", onInput SetTaskDuration ] []
        , button [ onClick (AddTask model.scratch) ] [ text "Add a task" ]
        , button [ onClick (RemoveTask model.scratch) ] [ text "Remove a task (by name)" ]
        , ul [] (List.map (\l -> li [] [ text l.name ]) model.tasks)
        ]
